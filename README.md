# :lock: Passager

> A simple, secure and personal password manager.

Some passager features:

- Easy to use. Manage your passwords locally using a simple command line interface (CLI).
- You own your data. Move your database file to wherever you want.
- Secure by default. Use Advanced Encryption Standard (AES) by default. You can set your own encryption key.

## Usage

### Generate

Generate a random and secure password:

```
$ passager generate
```

### Add

Add a new credential record:

```
$ passager add MyCredential my@email.com supersecret
```

### Get

Get an existing credential record:

```
$ passager get MyCredential
```

### List

List all records:

```
$ passager list
```

### Delete

Delete an existing record:

```
$ passager delete MyCredential
```

### Help

All commands have a helper:

```
$ passager --help
```

## Settings

It is possible to configure your `passager` environment by settings the following environment variables:

- `PASSAGER_DATABASE`: path to the database file. Default: `~/.passager.db`.
- `PASSAGER_ENCRYPTION_KEY`: default key to be used for data encryption. The encryption key should have 16 bytes.

## Development

### Build

Build the project using Makefile:

```
$ make build
```

### Tests

Run all tests and calculate the code coverage using Makefile:

```
$ make tests
```

## TODO

- [X] Improve GET using LIKE
- [X] Add version and -v command
- [ ] Allow user to set an encryption key using envvars
- [ ] Increase code coverage to at least 60%
- [ ] Improve README documentation
- [ ] Improve package documentation
- [ ] Add CONTRIBUTION guide
- [ ] Add project license
- [ ] Run tests on github actions
- [X] Allow user to set database path using envvar
