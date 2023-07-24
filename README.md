# i18n

i18n is a Go [package](#package-i18n) that helps you translate Go programs into several languages.

## Package i18n

The i18n package provides support for looking up messages according to a set of locale preferences.

```go
go get github.com/makhmudovazeez/i18n
```

Set up the location of your json translation files.

```go
i18n.Location = "example_location_folder/"
```

Set up the location of your application language.
```go
i18n.Language = "en"
```

Load translations.

```go
i18n.T("example.message")
```

## License

i18n is available under the MIT license. See the [LICENSE](LICENSE) file for more info.
