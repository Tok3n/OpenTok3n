OpenTok3n
=========

OpenTok3n is an open source platform for multi-factor authentication based on Tok3n Platform.

Requirements
-
* [Go programming language][1].
* [Git][2] and [Mercurial][3] for the `go get` dependencies resolver. See more information [about the version control tools][7].
* A MySQL server and a new database.

Install
-
1. Verify your Golang installation running `go`.
2. Create/Select a directory for the OpenTok3n installation and set up the $GOPATH in your bash profile running `echo "export GOPATH=$HOME/SelectedDirectory" >> $HOME/.bashrc`. For more information about choosing a GOPATH run `go help gopath`. Also see [this post][6] about choosing a bash profile file.
3. Install OpenTok3n running `go get github.com/Tok3n/OpenTok3n` and then `go install github.com/Tok3n/OpenTok3n`.
4. Verify the installation running `OpenTok3n`.

Create a *tok3n.com* account
-

1. Install [Tok3n][5] on your Android smartphone.
2. Create an account at [secure.tok3n.com][4]. Scan the barcode with the Tok3n app to authenticate.
3. Click the "New integration" button and follow the steps.
4. Retrieve your API Keys.


Configuration
-
There are three configuration methods:

* Perform an interactive configuration running `OpenTok3n` for the first time.
* Create an *opentok3n.config* file and provide the path to the folder where such file is located. `OpenTok3n --path /pathToConfigFolder`. Please note the file should be written in json format.

```json
{
	"Inited": true,
	"Address": "127.0.0.1",
	"Port": "63568",
	"DBAddress": "localhost",
	"DBPort": "3306",
	"DBUser": "root",
	"DBPassword": "yourDBpassword",
	"DBName": "opentok3n",
	"DBTablePrefix": "myintegration_",
	"Tok3nAPISecret": "YOUR-TOK3N-SECRET-KEY",
	"Tok3nAPIKey": "YOUR-TOK3N-PUBIC-KEY"
}
```

* Provide the configuration parameters using the "--" syntax to the server at run. Learn more running `OpenTok3n --help`.

Further considerations
-

The default values for `Address` and `Port` are "localhost:63568".

You should manually set `Inited` as true until we fix the auto file loader.

You can also run `OpenTok3n --help` if you forget the installation arguments.

[1]: http://golang.org/doc/install
[2]: http://git-scm.com/book/en/Getting-Started-Installing-Git
[3]: http://mercurial.selenic.com/wiki/Download
[4]: http://secure.tok3n.com
[5]: https://play.google.com/store/apps/details?id=com.secureyourself.with.tok3n
[6]: http://www.joshstaiger.org/archives/2005/07/bash_profile_vs.html
[7]: https://code.google.com/p/go-wiki/wiki/GoGetTools