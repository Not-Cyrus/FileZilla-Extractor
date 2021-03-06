# FileZilla-Extractor

### NOTE: Any other operating system but windows will ask you to specify the path.

Allows you to easily extract server login data from FileZilla (host/port/user/pass/keyfile (if it has one)) and store it in a JSON format for easy readability.

### Install
[Windows binary only](https://github.com/Not-Cyrus/FileZilla-Extractor/releases)


### Building from the source code

```bash
git clone https://github.com/Not-Cyrus/FileZilla-Extractor

cd FileZilla-Extractor

go get -v -t -d ./...

go build
```

### Run 

You can either double-click to run, or use the command line.

```
C:\Users\Cyrus\Desktop\FileZillaExtractor>.\FileZilla-Exporter.exe -h
NAME:
   FileZilla Exporter - Export server logins/key files

USAGE:
   main.exe [global options] [arguments...]

VERSION:
   0.0.1

GLOBAL OPTIONS:
   --savetype value, --st value      switch the save type between all/managed/recent (default: "all")
   --verbose, --vv                   verbose (default: false)
   --help, -h                        show help (default: false)
   --version, -v                     print the version (default: false)
```

```
C:\Users\User\Desktop\FileZillaExtractor>.\FileZilla-Exporter.exe --st all
Exported 1 managed servers
Exported 10 recently connected servers
```

```
C:\Users\User\Desktop\FileZillaExtractor>.\FileZilla-Exporter.exe --st managed
Exported 1 managed servers
```

```
C:\Users\User\Desktop\FileZillaExtractor>.\FileZilla-Exporter.exe --st recent
Exported 10 recently connected servers
```

