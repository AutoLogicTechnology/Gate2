# Gate2

Simple REST API exposing TOTP 2-Factor authentication over HTTP(S) for easy integration.

# Key Features

* Easily integrate TOTP into your application(s)
* Generates a QR code when a new user is created
* Also generates scratch codes as a backup system
* Supports three SQL database backends for storing users
 * MySQL
 * PostgreSQL
 * SQLite
* It's written in Go (1.3)
* GPLv3

# Installation

You will need Go 1.3. Check out http://golang.org for more information on installing Go.

1. Ensure your `$GOPATH` is configured
1. `go get github.com/AutoLogicTechnology/Gate2/cmd/gate2-api`

You will now have Gate2's API command in your `$GOPATH/bin` directory.

## Binaries & Packages

Both are coming soon. A custom repository will be setup offering Deb and RPM packages.

# Usage

Using the Gate2 API is very easy. You simply define your configuration file as such:

```json
{
    "database": {
        "engine": "sqlite3",
        "href": "./gates.sql",
        "purge": true
    }
}
```

And save the file to a location of your choosing. Then you launch the API: `./gate2-api -config "path/to/config.file"`

# Bugs

Please feel free to raise issues and report bugs. All are welcome.
