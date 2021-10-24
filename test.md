# Bump Version in Package.json & Release
Increments the version following semver conventions, updates the version in package.json and creates a release with a generated changelog.

## Inputs
| Name | Description | Required | Default |
| --- | --- | --- | --- |
| git-chglog-version | git-chglog version to install. Defaults to latest. | false | `latest` |
| go-version | Go version to install. Defaults to ^1.17. | false | `^1.17` |
| tag-prefix | Prefix to add to the created tag. | true | `` |
| project-name | The project scope this change relates to | true | `` |
| project-path | The directory to the package.json to update | false | `.` |
| changelog-config | Path to the configuration file for git-chglog. | false | `.chglog/releaselog-config.yaml` |

## Outputs
| Name | Description | Value |
| --- | --- | --- |
| new-tag | New tag created for the release. | `${{ steps.bump-version.outputs.new_tag }}` |

## External Actions
| Name | Creator | Version | Step Name | Step ID |
| --- | --- | --- | --- | --- |
| [setup-go](https://github.com/actions/setup-go/tree/v2) | actions | v2 | Set up go |  |
| [cache](https://github.com/actions/cache/tree/v2) | actions | v2 | Set up Go cache |  |
| [github-tag-action](https://github.com/mathieudutour/github-tag-action/tree/v5.6) | mathieudutour | v5.6 | Bump version | bump-version |
| [git-auto-commit-action](https://github.com/stefanzweifel/git-auto-commit-action/tree/v4) | stefanzweifel | v4 | Commit version | auto-commit-action |
| [release-action](https://github.com/ncipollo/release-action/tree/v1) | ncipollo | v1 | Create release |  |
