# Go WireGuard Public Key Prefixer

This project is a Go-based WireGuard public key prefixer, inspired by [wireguard-vanity-address](https://github.com/warner/wireguard-vanity-address). It generates WireGuard private and public keys until it finds a public key that starts with a specified prefix.

## Usage

To use this program, simply run it with the desired prefix as an argument:

```bash
go run main.go <prefix>
```

The program will generate keys and print the number of keys generated every second. When it finds a key with the specified prefix, it will print the private and public keys and exit.

## Limitations

The program only supports prefixes up to 10 characters long. The search is case-sensitive.

## License

This project is licensed under the MIT License.
