# Passager

A simple, secure and personal password manager.

## Build

```
$ go build -o passager cmd/passager/main.go
```

## Usage

### Generate

```
$ passager generate
```

### Add

```
$ passager add GMAIL my@email.com supersecret -k #encryption-key#
```

### Get

```
$ passager get GMAIL -k #encryption-key#
```

### List

```
$ passager list -k #encryption-key#
```

### Delete

```
$ passager delete GMAIL
```

## Notes

- The encryption key should have 16 bytes
