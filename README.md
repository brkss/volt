# Volt

encrypt files with AES using passkey.

## Installation

to install volt locally first clone this repository :

```bash
git clone https://github.com/brkss/volt
```

cd to volt's folder

```bash
cd volt
```

install using make

```bash
make install
```

## Usage

to encrypt a file :
```bash
volt e input.txt [password] output.txt 
```

to decrypt a file :
```bash
volt e encrypted-input.txt [password] > decrypted-output.txt 
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)