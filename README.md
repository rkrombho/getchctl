getchctl
========
`getchctl` is a CLI client for [Getch](https://github.com/rkrombho/getch).

### Features

The following Getch features are supported by `getchctl`
* [List values](https://github.com/rkrombho/getch/wiki/List-values) - `getchctl list`
* [Query values](https://github.com/rkrombho/getch/wiki/Query-values) - `getchctl get <key>`
* [Query files](https://github.com/rkrombho/getch/wiki/Query-files) - `getchctl getfile <filename>`
* [Files as templates](https://github.com/rkrombho/getch/wiki/Files-as-templates) - `getchctl getfile <filename>`
* [Encryption](https://github.com/rkrombho/getch/wiki/Encryption) - `getchctl encrypt <filename>`

### Configuration

The `getchctl` client expects to be configured thorugh environment variables:

```bash
# mandatory
export GETCH_SERVER="http://<getchhost>:<getchport>/getch"
# optional
export GETCH_CACERT="/path/to/CA/cert" # PEM encoded CA certs file to be used when running Getch on SSL
export GETCH_CLIENT_IP="192.125.33.16" # IP of the NIC you want to use. Should be DNS resolvable (on Getch server side)

```
