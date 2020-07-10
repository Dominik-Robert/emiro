# Available Arguments
The emiro CLI has many arguments you can pass in. Here I will show you all arguments with examples and all available parameters

## search

## exec

## new

## show

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