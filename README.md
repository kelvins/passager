# :lock: Passager

> A simple, secure and personal password manager.

Some `passager` features:

- Easy to use. Manage your passwords locally using a simple command line interface (CLI).
- You own your data. Move your database file to wherever you want.
- Secure by default. Use Advanced Encryption Standard (AES) by default. You can set your own encryption key.

## 🚸 Usage

`generate` a random and secure password:

```
$ passager generate
```

`add` a new credential:

```
$ passager add MyCredential my@email.com supersecret
```

`get` an existing credential:

```
$ passager get MyCredential
```

`list` all credentials:

```
$ passager list
```

`delete` an existing credential:

```
$ passager delete MyCredential
```

## ⚙️ Settings

It is possible to configure your `passager` environment by settings the following environment variables:

- `PASSAGER_DATABASE`: path to the database file. Default: `~/.passager.db`.
- `PASSAGER_ENCRYPTION_KEY`: data encryption key. The encryption key should have 16 bytes.

## 👷 Development

#### Build

Build the project using:

```
$ make build
```

Run all tests and calculate the code coverage using:

```
$ make tests
```

## TODO

- [X] Improve GET using LIKE
- [X] Add version and -v command
- [X] Allow user to set an encryption key using envvars
- [ ] Increase code coverage to at least 60%
- [X] Improve README documentation
- [ ] Improve package documentation
- [ ] Add CONTRIBUTION guide
- [ ] Add project license
- [ ] Run tests on github actions
- [X] Allow user to set database path using envvar
