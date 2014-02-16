OpenTok3n
=========

We are an Open Source platform for multi-factor authentication. Based in the Tok3n Platform.

Requirements
-
* [Go programming language][1]
* [Git Client app][2]
* [Mercurial client app][3]
* A MySQL database

Installing
-
* Once you downloaded the [GoLang][1]. (Ensure that you have configured a $GOPATH)
* Add $GOPATH/bin to the $PATH
* Run the following commands.

go get github.com/Tok3n/OpenTok3n

* That's it the OpenTok3n App is now installed in your machine.

Configuration
-

Fist you need to have an integration with [Tok3n][4]. For you to have the API Keys that are going to syncronize the end user smartphones with your integration. (The tok3n platform provides the high entropy random number generators for the integration user keys)

To configure OpenTok3n you can do it in three ways:

* Provide a configuration file to the server at run.
* Provide the parameters via the " - " sintax to the server at run.
* Runing the server and perform an interactive configuration. 

The configuration file is a json string encoded that should looks like the following:
```json
{
	"Inited":true,
	"Address":"127.0.0.1",
	"Port":"63568",
	"DBAddress":"localhost",
	"DBPort":"3306",
	"DBUser":"root",
	"DBPassword":"root",
	"DBName":"open",
	"DBTablePrefix":"open_",
	"Tok3nAPISecret":"THIS-IS-YOUR-SECRET-KEY",
	"Tok3nAPIKey":"THIS-IS-YOUR-PUBIC-KEY"
}
```

You can always run _OpenTok3n --help_ to remember this detils
Now you can run the command  to view the parameters that OpenTok3n acept 

[1]: http://golang.org/doc/install
[2]: http://git-scm.com/book/en/Getting-Started-Installing-Git
[3]: http://mercurial.selenic.com/wiki/Download
[4]: http://secure.tok3n.com