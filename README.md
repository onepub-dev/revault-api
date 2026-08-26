# reVault API bindings

reVault is an encrypted archive and local-vault library for files, credentials, keys, and typed records. This companion repository publishes the Go module, Swift package, and PHP Composer package from the same versioned native ABI. See the [reVault documentation](https://github.com/onepub-dev/reVault/tree/main/docs).

## Go

```sh
go get github.com/onepub-dev/revault-api@v0.3.12
```

## Swift

```swift
.package(url: "https://github.com/onepub-dev/revault-api", exact: "0.3.12")
```

## PHP

```sh
composer require onepub/revault-api:0.3.12
```

All three packages use the same ABI-v3 native artifacts. Secret variable and form-field reads are callback-scoped and their temporary native storage is wiped after use.
