# Available Arguments
The emiro CLI has many arguments you can pass in. Here I will show you all arguments with examples and all available parameters


## search
Search is one of the important features. With search you can query against all saved commands on a server and get a detailed list of all commands or the commands that match your query.

### Usage
emiro search QUERY [flags] | emiro search -a [flags]


### Examples

```bash
# Query against all
emiro search -a

# Query against all that match the QUERY without whitespaces:
emiro search QUERY

# Query against all that match the QUERY with whitespaces:
emiro search "MY QUERY"

# Limit the output of the search command
```


### Options
```
  -a, --all             Shows all existing entries
  -c, --count int32     Sets the maximum count of entries (default 10)
  -h, --help            help for search
  -o, --output string   Set the output format. Options are: short, wide, json, detailed (default "short")
```

### Options inherited from parent commands
```
      --aliasFile string         Specify the aliasFile where the alias commands are stored to (default "$(HOME)/.profile")
      --config string            config file (default is $PWD/emiro.yaml)
      --databaseHost string      Set the database host (default "http://localhost")
      --databaseIndex string     Sets the index which are used (default "emiro")
      --databaseInsecure         enable insecure database connection
      --databasePort int         Set database port (default 9200)
      --databaseType string      Set the database Type  (default "elasticsearch")
      --emiroHost string         specify the emiro host (default "104.197.10.159")
      --emiroPort int            specify the emiro host port (default 9000)
      --emiroServerHost string   Specify the listening host for the server (default ":")
      --emiroServerPort int      Specify the listening port for the server (default 9000)
      --tempDir string           Specifies the tempDir (default "/tmp/emiro")
  -v, --verbose                  Activates the verbose output
```


## exec
Executes a snippet from the terminal. You can write a query which results in one command (this will automatically executed) 
or you can refer to an id with --id where you don't have to write a more detailed query.

### Usage
emiro exec [flags]


### Examples


### Options

```
      --alias string        Specify if the command will create an alias in your system
      --append string       Append the argument after the command
  -h, --help                help for exec
  -p, --param stringArray   Specify a Parameter array to change the command
      --prepend string      Prepend the argument before the command
      --remote string       Specify the host where the command will run
```


### Options inherited from parent commands
```
  --aliasFile string         Specify the aliasFile where the alias commands are stored to (default "$(HOME)/.profile")
  --config string            config file (default is $PWD/emiro.yaml)
  --databaseHost string      Set the database host (default "http://localhost")
  --databaseIndex string     Sets the index which are used (default "emiro")
  --databaseInsecure         enable insecure database connection
  --databasePort int         Set database port (default 9200)
  --databaseType string      Set the database Type  (default "elasticsearch")
  --emiroHost string         specify the emiro host (default "104.197.10.159")
  --emiroPort int            specify the emiro host port (default 9000)
  --emiroServerHost string   Specify the listening host for the server (default ":")
  --emiroServerPort int      Specify the listening port for the server (default 9000)
  --tempDir string           Specifies the tempDir (default "/tmp/emiro")
  -v, --verbose                  Activates the verbose output
```

## new
New creates a new command. Actually the only option is to import the new command from a json file. 
Specify it within an editor or per Line from the tool is on the roadmap.

### Usage
emiro new --fromFile YOUR_FILE.json

### Examples

### Options

```
  -d, --description string   Help message for toggle
      --fromFile string      Specify a file that is used
  -h, --help                 help for new
  -i, --interactive          Help message for toggle
  -n, --name string          Help message for toggle
```

### Options inherited from parent commands

```
      --aliasFile string         Specify the aliasFile where the alias commands are stored to (default "$(HOME)/.profile")
      --config string            config file (default is $PWD/emiro.yaml)
      --databaseHost string      Set the database host (default "http://localhost")
      --databaseIndex string     Sets the index which are used (default "emiro")
      --databaseInsecure         enable insecure database connection
      --databasePort int         Set database port (default 9200)
      --databaseType string      Set the database Type  (default "elasticsearch")
      --emiroHost string         specify the emiro host (default "104.197.10.159")
      --emiroPort int            specify the emiro host port (default 9000)
      --emiroServerHost string   Specify the listening host for the server (default ":")
      --emiroServerPort int      Specify the listening port for the server (default 9000)
      --tempDir string           Specifies the tempDir (default "/tmp/emiro")
  -v, --verbose                  Activates the verbose output
```

## show
Get detailed information about a command

### Usage
emiro show QUERY [flags]


### Examples
```bash
	emiro show Hostname

	# results in following output
	
	ID            hveX_3IB8mo66xNibt7z
	Name          echo Hostname
	Description   Prints the Hostname of the machine
	Command       echo $HOSTNAME
	Path          /bin/bash
	Language      
	OS            []
	Script        false
	Params				null
```
### Options

```
  -h, --help    help for show
  -p, --plain   Shows the result in plain json
```

### Options inherited from parent commands

```
      --aliasFile string         Specify the aliasFile where the alias commands are stored to (default "$(HOME)/.profile")
      --config string            config file (default is $PWD/emiro.yaml)
      --databaseHost string      Set the database host (default "http://localhost")
      --databaseIndex string     Sets the index which are used (default "emiro")
      --databaseInsecure         enable insecure database connection
      --databasePort int         Set database port (default 9200)
      --databaseType string      Set the database Type  (default "elasticsearch")
      --emiroHost string         specify the emiro host (default "104.197.10.159")
      --emiroPort int            specify the emiro host port (default 9000)
      --emiroServerHost string   Specify the listening host for the server (default ":")
      --emiroServerPort int      Specify the listening port for the server (default 9000)
      --tempDir string           Specifies the tempDir (default "/tmp/emiro")
  -v, --verbose                  Activates the verbose output
```

## delete


## server

## remote
Remote starts a remote session


## version
Version shows you the current version you use. 

Example
```bash
emiro version
```

Results in following output

```bash
# Output
Version:  0.9.0
```

Version has no other parameter

## help

Help prints a detailed help message

Example
```bash
emiro help
```
Results in following output

```bash
# Output
Emiro was build for the problem that you need to leave your terminal for every snippet you want to insert in your terminal. 
With Emiro you can search for specific snippets and execute them. 
If you don't have the needed dependencies you can even run it everywhere(where you ran emiro remote service is running.

Usage:
  emiro [command]

Available Commands:
  delete      Deletes no more needed commands
  exec        Execute a snippet
  help        Help about any command
  new         Creates a new Command
  remote      Runs a remote server
  search      Search and List the available command
  server      Starts an emiro server
  show        Get detailed information about a command
  version     Prints the current version

Flags:
      --aliasFile string         Specify the aliasFile where the alias commands are stored to (default "$(HOME)/.profile")
      --config string            config file (default is $PWD/emiro.yaml)
      --databaseHost string      Set the database host (default "localhost")
      --databaseIndex string     Sets the index which are used (default "emiro")
      --databaseInsecure         enable insecure database connection
      --databasePort int         Set database port (default 9200)
      --databaseType string      Set the database Type  (default "elasticsearch")
      --emiroHost string         specify the emiro host (default "localhost")
      --emiroPort int            specify the emiro host port (default 9000)
      --emiroServerHost string   Specify the listening host for the server (default ":")
      --emiroServerPort int      Specify the listening port for the server (default 9000)
  -h, --help                     help for emiro
      --tempDir string           Specifies the tempDir (default "/tmp/emiro")
  -v, --verbose                  Activates the verbose output

Use "emiro [command] --help" for more information about a command.
```

You can call the help function on every command

For example you can call the help function for exec:

```bash
emiro help exec
```

Will result into following output

```bash
Executes a snippet from the terminal. You can write a query which results in one command (this will automatically executed) 
or you can refer to an id with --id where you don't have to write a more detailed query.

For example:

emiro exec "echo hostname"

# Will execute the query and execute the resulting command.

You can execute your snippet on a remote machine with --remote HOST:PORT.

You can append or prepend arguments to your command with --prepend or --append, if you write more arguments than one (first is the query) 
then the others will automatically appended to the command.

Usage:
  emiro exec [flags]

Aliases:
  exec, run

Flags:
      --alias string        Specify if the command will create an alias in your system
      --append string       Append the argument after the command
  -h, --help                help for exec
  -p, --param stringArray   Specify a Parameter array to change the command
      --prepend string      Prepend the argument before the command
      --remote string       Specify the host where the command will run

Global Flags:
      --aliasFile string         Specify the aliasFile where the alias commands are stored to (default "$(HOME)/.profile")
      --config string            config file (default is $PWD/emiro.yaml)
      --databaseHost string      Set the database host (default "localhost")
      --databaseIndex string     Sets the index which are used (default "emiro")
      --databaseInsecure         enable insecure database connection
      --databasePort int         Set database port (default 9200)
      --databaseType string      Set the database Type  (default "elasticsearch")
      --emiroHost string         specify the emiro host (default "localhost")
      --emiroPort int            specify the emiro host port (default 9000)
      --emiroServerHost string   Specify the listening host for the server (default ":")
      --emiroServerPort int      Specify the listening port for the server (default 9000)
      --tempDir string           Specifies the tempDir (default "/tmp/emiro")
  -v, --verbose                  Activates the verbose output
```