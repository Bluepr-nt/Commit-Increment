# Commit-Increment
Commit-Increment is a Go library and a GitHub Action that identifies the increment level from conventional commit messages.

## GitHub Action

The Commit Increment Action determines the version increment level based on commit messages. It uses patterns for major and minor version increments, which can be customized.

### Inputs

- `major_pattern`: Pattern for major version increment. Default is 'MAJOR:'.
- `minor_pattern`: Pattern for minor version increment. Default is 'MINOR:'.
- `install_go`: Whether to install Go. Default is 'true'.

### Outputs

- `increment_level`: The determined version increment level.
- Possible ouputs: `major`, `minor` and defaults to `patch` when there is no match

### Example usage

```yaml
steps:
  - name: Determine version increment level
    id: increment
    uses: your-github-username/commit-increment@v1
    with:
      major_pattern: 'Major:'
      minor_pattern: 'Minor:'
      install_go: 'true'
  - name: Use increment level
    run: echo "The increment level is ${{ steps.increment.outputs.increment_level }}"
```

### Suggested pattern set
#### Opinionated [Conventional Commit](https://conventionalcommits.org/en/v1.0.0/)
```yaml
major_pattern: "^((build|ci|docs|feat|fix|perf|refactor|test)(\([a-z 0-9,.\-]+\))?!: [\w \(\),:.;\-#&']+|\nBREAKING CHANGES: [\w \(\),:.;\-#&']+)$"  
# See https://regex101.com/r/ORB9yp/1
minor_pattern: "^(feat)(\([a-z 0-9,.\-]+\))?!?: [\w \(\),:.;\-#&']+$"  
# See https://regex101.com/r/pBspGO/1
```

#### Simplified [Conventional Commit](https://conventionalcommits.org/en/v1.0.0/)
```yaml
major_pattern: ".*!:.*"  
minor_pattern: "^feat:.*:"  
```

#### Simple pattern
```yaml
major_pattern: "Major:"  
minor_pattern: "Minor:"  
```
