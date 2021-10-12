# Passager

A simple, secure and personal password manager.

## Build

```console
$ go build -o passager cmd/passager/main.go
```

## Usage

### Generate

```
$ passager generate
```

### Add

```
$ passager add GMAIL my@email.com supersecret -k my-16-bytes-key
```

### Get

```
$ passager get GMAIL -k my-16-bytes-key
```

### List

```
$ passager list -k my-16-bytes-key
```

### Delete

```
$ passager delete GMAIL
```
