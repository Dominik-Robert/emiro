# Emiro - Search, find and execute shell snippets without leaving your terminal

Emiro was build for the problem that you need to leave your terminal for every snippet you want to insert in your terminal. With Emiro you can search for specific snippets and execute them. If you don't have the needed dependencies you can even run it everywhere(where you ran emiro remote service is running.

## Documentation
A complete and more detailed documentation is on:
[Emiro documentation](https://github.io/dominik-robert/emiro)

## Install 

Before emiro is not available in the package manager you have to download it here from the release page.
[Emiro Release Page](https://github.com/Dominik-Robert/emiro/releases/tag/) and click on the current version then Download for your OS


```wget
# Downlad emiro for linux and put it in the bin directory. Change before the VERSION string to VERSION
wget https://github.com/Dominik-Robert/emiro/files/4880723/emiro-linux-$(VERSION).zip 
unzip emiro-linux-v-VERSION.zip 
chmod +x emiro-linux && mv emiro-linux /usr/local/bin/emiro

# Downlad emiro for osx and put it in the bin directory. Change before the VERSION string to VERSION
wget https://github.com/Dominik-Robert/emiro/files/4880723/emiro-darwin-${VERSION}.zip 
unzip emiro-darwin-v-VERSION.zip 
chmod +x emiro-darwin && mv emiro-darwin /usr/local/bin/emiro

# For windows there is no existing shortcut. So please go to the release page and download the binary for windows 
```

## Usage or Quickstart
You can use the emiro CLI program on windows,linux and osx. 
Please install the CLI with the explanation on this page.

### search
You can search through all existing commands with the search argument.

```bash
# Lists all commands that match the query
emiro search "QUERY" 

# for the example below try this:
emiro search "echo Hostname"

# or take a detailed look with different output:

emiro search "echo Hostname" -o detailed
emiro search "echo Hostname" -o short # default
emiro search "echo Hostname" -o wide
emiro search "echo Hostname" -o json

```

### exec
You can exec all existing commands that match your OS. For example on linux and osx this should be done successfull
```bash
emiro exec "echo Hostname"
# will exec your hostname
```

## Available Commands
To get a list of the existing commands and arguments please visit the command page:

[List of available commands](https://github.io/dominik-robert/emiro/commands.html)


## How does emiro work?
Emiro works with a server. So the client (cli) will connect with a grpc call to a emiro-server. The server will search in a database (offical is elasticsearch) for the matching queries and returns them to the client.

## Contribute

Feel free to contribute. Everything is welcome from PullRequests to FeatureRequests or to find new issues.


## Roadmap

* [ ] Save the script local
  * It would be great to implement a save option to run snippets with specific parameter to keep dynamic 
* [ ] edit
  * It should be possible to edit the snippets from the terminal
* [ ] access rights
  * you can execute the snippets on a different server, so it would be important to limit the execution rights.
* [ ] More detailed search
  * The search must be more detailed, so that you can view directly the command






