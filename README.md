# scopeCheck
Check Urls for domain matching passed via arguments.


## Install
```
go install github.com/shubhack-repo/scopeCheck@latest
```

## Usage
```
cat urls.txt
http://kali.com/path1
https://kali.com/path2
http://abc.kali.com/path3
http://anything.com/path4
```

### Output

```
cat urls.txt | scopeCheck -d kali.com
http://kali.com/path1
https://kali.com/path2
http://abc.kali.com/path3
```
