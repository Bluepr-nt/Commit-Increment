# Commit-Increment
Commit-Increment is a GitHub Action that identifies the increment level from conventional commit messages. It can also be used as a standalone Go library.

## GitHub Action

The Commit Increment Action determines the version increment level based on commit messages. It uses patterns for major and minor version increments, which can be customized.

### Inputs

- `major_pattern`: Pattern for major version increment. Default is 'MAJOR:'.
- `minor_pattern`: Pattern for minor version increment. Default is 'MINOR:'.
- `install_go`: Whether to install Go. Default is 'true'.

### Outputs

- `increment_level`: The determined version increment level.

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
Conventional Commit (link)
  Choose your desired keywords
Simplified
Complex

