# :lock: Passager

> A simple, secure and personal password manager.

Some `passager` features:

- Easy to use. Manage your passwords locally using a simple command line interface (CLI).
- You own your data. Move your database file to wherever you want.
- Secure by default. Use Advanced Encryption Standard (AES) by default.

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
```

## 🚸 Usage

`generate` a random and secure password:

```
$ passager generate
```

`add` a new credential:

```
$ passager add MyCredential my@email.com supersecret
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

When in doubt, use the `help` flag:

```
$ passager add -h
```

## 🛠️ Settings

It is possible to configure your `passager` environment by settings the following environment variables:

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
