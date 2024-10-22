## 0.1.5

Website:
- Created new website (https://www.promptorium.org) for documentation
- Updated CI/CD pipeline to automatically update the website on every release
- Added new documentation page dedicated to configuration

Configuration:
- Added `promptorium init` command to initialize config and theme files. It is recommended to run this command after installing promptorium.

Installation:
- Improved installation instructions
- Added script for manual installation(without a package manager)

Fixes:
- Fixed separator being printed when there are no components on the right side
- Added fallback to default config if config file is not found (same for theme file)
- Fixed theme's component dividers being ignored.
- Improved deb package installation, added `scripts/deb/conffiles`

Other changes:
- Changed `--config` flag to `--config-file`
- Changed `--theme` flag to `--theme-file`