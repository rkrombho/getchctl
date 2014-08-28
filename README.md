getchctl
========
`getchctl` is a command line client for [Getch](https://github.com/rkrombho/getch).

### Features

The following Getch features are supported by `getchctl`
* [List values](https://github.com/rkrombho/getch/wiki/List-values) - `getchctl list`
* [Query values](https://github.com/rkrombho/getch/wiki/Query-values) - `getchctl get <key>`
* [Query files](https://github.com/rkrombho/getch/wiki/Query-files) - `getchctl getfile <filename>`
* [Files as templates](https://github.com/rkrombho/getch/wiki/Files-as-templates) - `getchctl getfile <filename>`
* [Encryption](https://github.com/rkrombho/getch/wiki/Encryption) - `getchctl encrypt <filename>`

### Download and Install
Just download the [latest release](https://github.com/rkrombho/getchctl/releases/) for your platform and place it anywhere on your system.
It's of course recommendable to extend your `$PATH` or `%PATH%` variable to the location where you plan to locate the `getchctl` or `getchctl.exe` file.

### Configuration

The `getchctl` client expects to be configured thorugh environment variables:

```bash
# mandatory
export GETCH_SERVER="http://<getchhost>:<getchport>/getch"
# optional
export GETCH_CACERT="/path/to/CA/cert" # PEM encoded CA certs file to be used when running Getch on SSL
```

### Usage
```bash
NAME:
   getchctl - A command line client for Getch

USAGE:
   getchctl [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   get		retrieve a value for a given key
   list		List all values stored in Getch for the host where this command is executed from
   getfile	Downloads the file with the given name from Getch
   encrypt	encrypts a given value so that it can be used in Getch server-side configurations
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version

```
