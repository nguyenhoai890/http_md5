## Summary
This is a tool that makes http requests and prints the address of the request along
with the MD5 hash of the response.

If there is any error, the tool will print an error message next to the url name instead of the MD5 hash.

## Basics
To be able to work on this repo, you must have Go installed on your machine.
Please follow the directions for installing Go on the [Go Website](https://go.dev/doc/install).

* Build: you can use `make build`. The output ```myhttp``` binary will be placed in ```./build``` by default.
* Test:  you can use ```make test```.

## Usage
### Command:
```bash 
$ ./build/myhttp [-parallel <parallel>] <urls>
```
### Parameter:
* `urls`: If there is any domain that does not have a scheme, the tool will auto set Scheme to `http`. For example: `google.com` will be processed as `http://google.com`

### Flag:
* `-parallel`: (Optional) Set the number of go routines that can run at the same time. The default value is `10` which will be used when the flag is not provided or passed flag has a value less than `1`. 

### Example:
```bash
$ ./build/myhttp google.com
$ ./build/myhttp -parallel 2 google.com yahoo.com reddit.com
```