# :lock: Passager

A simple, secure and personal password manager.

## Build

Build the project:

```
$ make build
```

## Tests

Run the tests:

```
$ make tests
```

## 🚸 Usage

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

## :shield: Encryption

Use Advanced Encryption Standard (AES) encryption algorithm.

## 📓 Notes

- The encryption key should have 16 bytes

## TODO

- [X] Improve GET using LIKE
- [X] Add version and -v command
- [ ] Increase code coverage to at least 60%
- [ ] Improve README documentation
- [ ] Improve package documentation
- [ ] Add CONTRIBUTION guide
- [ ] Add project license
- [ ] Run tests on github actions
- [ ] Allow user to set database path using envvar
