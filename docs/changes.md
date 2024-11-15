## 0.1.6

Breaking changes:
- Renamed `git_color_no_branch` to `git_color_no_repository` in the config file
- Renamed `git_color_no_remote` to `git_color_no_upstream` in the config file

Performance improvements:
- Added lazy loading and caching of all context data (git status, os, shell, etc.)
- Improved git status data retrieval

Overall improvements:
- Added completion for bash and zsh
- Improved debug logging

Refactorings:
- Added `context` package to confmodule, which is responsible for retrieving context data
- Refactored the way modules are loaded and executed, this should make it easier to add new modules and potentially load them on the fly

Fixed bugs:
- Fixed a bug where the exit code was not being correctly set
