# Crypter 🔐

**Crypter** is a simple, secure command-line utility written in Go for encrypting and decrypting files using a user-provided secret key. It is designed to be minimal, reliable, and easy to integrate into scripts or automation pipelines.

---

## 🚀 Installation

You can install Crypter using `go install`:

```bash
go install github.com/mhthrh/crypto/crypter@latest
```

Make sure your `$GOBIN` is in your `PATH` to run the command directly:

```bash
export PATH=$PATH:$(go env GOBIN)
```

---

## 🔐 Features

* 🔒 AES-based file encryption with user-provided key
* 🔓 Decrypts previously encrypted files with the same key
* 🗘️ Output file has the same name as the original but **without the extension**
* 🧹 Simple CLI interface
* 🛡️ Designed for safety — files are not overwritten

---

## 🛠️ Usage

### 🔸 Encrypt a File

```bash
crypter encrypt --key YOUR_SECRET_KEY --path /path/to/yourfile.txt
```

* Creates: `/path/to/yourfile` (without `.txt`)
* Original file remains unchanged.

### 🔸 Decrypt a File

```bash
crypter decrypt --key YOUR_SECRET_KEY --path /path/to/yourfile
```

* Assumes the file at `/path/to/yourfile` is encrypted.
* Creates the decrypted file: `/path/to/yourfile.txt` (original extension may not be restored unless stored or known).

---

## 🧪 Example

```bash
# Encrypt
crypter encrypt --key mysupersecret --path ~/documents/report.pdf

# This will create: ~/documents/report

# Decrypt
crypter decrypt --key mysupersecret --path ~/documents/report
```

---

## ⚠️ Important Notes

* The encryption key (`--key`) is required and must be **kept safe**. If lost, encrypted files cannot be recovered.
* The application **does not delete or modify** the original file.
* Extensions are stripped from the encrypted file name; you are responsible for tracking the original extension if needed during decryption.

---

## 🔧 Flags

| Flag      | Type   | Description                                        |
| --------- | ------ | -------------------------------------------------- |
| `--key`   | string | **Required**. Secret key for encryption/decryption |
| `--path`  | string | **Required**. Path to the file to encrypt/decrypt  |
| `encrypt` |        | Command to encrypt the file                        |
| `decrypt` |        | Command to decrypt the file                        |

---

## 📆 Build from Source

Clone and build manually if desired:

```bash
git clone https://github.com/mhthrh/crypto.git
cd crypto/crypter
go build -o crypter
```

---

## 🔐 Security

* Files are encrypted using standard algorithms (e.g., AES-256 in GCM mode — specify in code/docs if applicable).
* Secret key should be strong (use at least 16 characters or more).
* Consider integrating with password managers or vaults for secure key management.

---

## 📄 License

MIT License © [mhthrh](https://github.com/mhthrh)

---

## 🤛🏻 Contributing

Feel free to open issues or pull requests! Feedback, improvements, and contributions are welcome.
