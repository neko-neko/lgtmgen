# lgtmgen
lgtmgen is LGTM image generator.  
Too easy to convert your favorite pictures.  

**Before**  
![](https://cloud.githubusercontent.com/assets/6947393/14233257/690a0228-f9fe-11e5-8d2f-5ece3b7b3ba9.jpg)

**After**  
![](https://cloud.githubusercontent.com/assets/6947393/14233259/7af4a416-f9fe-11e5-811b-974523a442e7.jpg)

## Installation
**[Download latest binary](https://github.com/neko-neko/lgtmgen/releases/latest)**

Multi-platform support
- Windows
- Mac OS X
- Linux 32,64bit

Or get the library
```
$ go get github.com/neko-neko/lgtmgen
```
Or clone the repository and run
```
$ go install github.com/neko-neko/lgtmgen
```

## Usage
```
Usage of lgtmgen:
  -d string
    	Input directory path(Short)
  -directory string
    	Input directory path
  -f	Force overwrite if output file exists(Short)
  -force
    	Force overwrite if output file exists
  -o string
    	Output directory path(Short)
  -output string
    	Output directory path
  -version
    	Print version information and quit.
```
### Example
```
$ lgtmgen -d /path/to/images/ -o /path/to/lgtms/
```

## Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## Credits
neko-neko

## License
MIT
