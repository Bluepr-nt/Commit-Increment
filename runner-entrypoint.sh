#!/usr/bin/env bash
set -euo pipefail

if [[ -z "${GITHUB_ORG:-}" ]]; then
  echo "GITHUB_ORG is required"
  exit 1
fi

RUNNER_SCOPE_URL="https://github.com/${GITHUB_ORG}"

fetch_org_registration_token() {
  if [[ -z "${GH_PAT:-}" ]]; then
    echo "GH_PAT is required when RUNNER_TOKEN is not provided"
    exit 1
  fi

  curl -fsSL -X POST \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${GH_PAT}" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    "https://api.github.com/orgs/${GITHUB_ORG}/actions/runners/registration-token" | jq -r '.token'
}

cd /home/runner/actions-runner

registration_token="${RUNNER_TOKEN:-${GITHUB_RUNNER_TOKEN:-}}"
if [[ -z "${registration_token}" ]]; then
  registration_token="$(fetch_org_registration_token)"
fi

if [[ -z "${registration_token}" || "${registration_token}" == "null" ]]; then
  echo "Failed to fetch runner registration token"
  exit 1
fi

if [[ ! -f .runner ]]; then
  config_args=(
    --unattended
    --url "${RUNNER_SCOPE_URL}"
    --token "${registration_token}"
  )

  if [[ -n "${RUNNER_GROUP:-}" ]]; then
    config_args+=(--runnergroup "${RUNNER_GROUP}")
  fi

  ./config.sh "${config_args[@]}"
else
  echo "Runner already configured, skipping config.sh"
fi

./run.sh
