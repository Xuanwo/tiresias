# Change Log

All notable changes to tiresias will be documented in this file.

## [v0.3.1] - 2017-07-22

### Fixed

- Fix config flag not handled correctly

## [v0.3.0] - 2018-07-20

### Added

- source: Check if consul reachable before connect

### Changed

- Save and load servers from level db in case of source is unavailable
- Skip wrong or unavailable source
- Auto clean not configured source and servers
- Show help message if no config passed

## [v0.2.0] - 2018-06-11

### Added

- source: Add consul support
- source: Add default settings

### Fixed

- Fix multi source support

## [v0.1.0] - 2018-06-05

### Added

- source: Add glob file path support

### Fixed

- utils: Fix seek position calculated incorrectly

## v0.0.1 - 2018-05-22

### Added

- Hello, tiresias!

[v0.3.1]: https://github.com/Xuanwo/tiresias/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/Xuanwo/tiresias/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/Xuanwo/tiresias/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/Xuanwo/tiresias/compare/v0.0.1...v0.1.0
