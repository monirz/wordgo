# wordgo  

![Wordgo Dictionary Demo](demo/demo.gif)


Wordgo is a cross platform desktop application (English to Bengali dictionary by default) written using Go QML binding. The app uses [gotri](https://github.com/monirz/gotri) for loading the words in the Trie/Prefix tree  and [therecipe/qt](https://github.com/therecipe/qt) a QT Go binding library for making the UI. 

For using it with different language you will need to replace the dictionary.txt file with the word-list of your language.    

This dictionary file has 135000 English word with multiple Bangla word menaing. Feel free to add any correction or add a new word to it. Also if you are good at QML and think you can improve the UI, don't hesitate, just jump into it. 

### Usage 
To use it just download the `wordgo` binary and the dictionary.txt from the deploy directory for your specific platform and run it.  
Run it from the root directory of the project.

```
$ ./deploy/linux/wordgo 
``` 

** Note: ** 
Currently you will need to have the dictionary.txt file in the same directory where you run the `wordgo` binary. So for the above cpommand you'll need to have the dictionary file in the root directory.    



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

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Make changes and add them (`git add .`)
4. Commit your changes (`git commit -m 'Added some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create new pull request 

## Credit 

The word-list is generated based on Ankur.org's Bangla-English [dictionary database](http://www.bengalinux.org/english-to-bengali-dictionary/about.htm.    






 