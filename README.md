# engo 

A cross platform English to Bengali dictionary written in Go. Engo is built using [gotri](https://github.com/monirz/gotri) and [therecipe/qt](https://github.com/therecipe/qt) a QT Go binding library for the UI. 

### Usage 
To use it just download the engo binary from the deploy directory for your specific platform and run it.  

## Build 
To install the dependency follow the [therecipe/qt](https://github.com/therecipe/qt) instruction for compiling QT application.  

```
$ make build 
```

The compiled binary file will be in the deploy/your-platform directory    
** compiling for Windows on Linux  ** 

```
$ make linux-to-windows 
```




 