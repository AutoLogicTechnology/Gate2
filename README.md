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

# Execution

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

# Usage & JSON 

The API has been designed to offer these endpoints:

* GET: `/totp/:id/:code`
* POST: `/totp/:id`
* DELETE: `/totp/:id`

The PUT HTTP verb is not included as the developer is expected to delete the target user and re-create their entry.

## :id & :code

The `:id` field can be basically anything that matches the regular expression `^[a-zA-Z0-9._@-]+$`, and `:code` must match `^[0-9]{6,8}$`.

## Responses

All responses are returned as JSON. XML is not an option. An example of POSTing to `/totp/:id` can be seen below. Not that the `qrcode` field has been left truncated due to the size of a Base63 encoded PNG file.

```json
{"message":"User added to the database successfully.","qrcode":"...","scratchcodes":["18036472","16892073","91460278"]}
```

This structure is very likely to change in the near future.

# Bugs

Please feel free to raise issues and report bugs. All are welcome.
