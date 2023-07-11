# go-nameb

[![GoDev][godev-image]][godev-url]

go-nameb パッケージは識別子名の命名規則変換を提供します.

## Usage

```go
nameb.Snake("SnakeCase") == "snake_case"
nameb.Kebab("kebab_case") == "kebab-case"
nameb.Camel("camel-case") == "camelCase"
nameb.Pascal("pascalCase") == "PascalCase"
```

## License

This software is released under the MIT License, see LICENSE.

## Author

17e10

[godev-image]: https://pkg.go.dev/badge/github.com/17e10/go-nameb
[godev-url]: https://pkg.go.dev/github.com/17e10/go-nameb
