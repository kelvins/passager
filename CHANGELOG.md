# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add `edit` command responsible for editing existing credentials.

### Changed

- Change the `add` command to make login optional.

### Fixed

- Show an error message when trying to delete an invalid credential.

## [0.2.0] - 2021-10-28

### Added

- Add this CHANGELOG to keep track of changes.

### Changed

- Change some command outputs to include a line break at the end.
- Sort credentials by name when listing.
- Print a specific message when there is no credentials to list.

### Fixed

- Perform a hard delete instead of a soft delete to avoid issues related to the unique name.

## [0.1.0] - 2021-10-19

### Added

- Create the commands: `generate`, `add`, `get`, `list` and `delete`.
- Use AES for password encryption.
- Allow users to configure the database path using environment variable.
- Allow users to configure the encryption key using environment variable.
