# :lock: Passager

> A simple, secure and personal password manager.

Some `passager` features:

- Easy to use: manage your passwords locally using a simple command line interface (CLI).
- Own your data: move your encrypted database file to wherever you want.
- Secure by default: use advanced encryption standard (AES) by default.

## :package: Installation

Download the appropriate binary for your OS and architecture from the [releases](https://github.com/kelvins/passager/releases) page, for example:

```
$ wget --no-check-certificate https://github.com/kelvins/passager/releases/download/v0.2.0/passager-linux-amd64
```

Make sure the file is executable:

```
$ chmod +x ./passager-linux-amd64
```

Move the file to a bin directory that is included in your PATH, for example:

```
$ sudo mv ./passager-linux-amd64 /usr/local/bin/passager
```

Check the `passager` version:

```
$ passager --version
Passager v0.2.0
```

## 🚸 Usage

`generate` a random and secure password:

```
$ passager generate
```

`add` a new credential:

```
$ passager add MyCredential -l my@email.com -p SuperSecret
```

> ⚠️ Use a whitespace prefix to your command to prevent it from being included in the history.

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

When in doubt, use the `help` flag, for example:

```
$ passager add -h
```

## 🛠️ Settings

It is possible to configure your `passager` environment by setting the following environment variables:

- `PASSAGER_DATABASE`: path to the database file. Default: `~/.passager.db`.
- `PASSAGER_ENCRYPTION_KEY`: data encryption key. The encryption key must be 16 bytes.

## 👷 Development

Build the project using:

```
$ make build-dev
```

Run all tests and calculate the code coverage using:

```
$ make tests
```
