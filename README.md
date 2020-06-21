# sni-scanner
Fastest Server Name Indication(SNI) proxy scanner.

## Notice
* Scan will use lot's of Network resource, setting a suitable thread num is important.
* most of VPS / Server provider monitor scanner behavior.

## Usage
```
go build
echo '1.1.1.1/24' > ip.txt
go run main.go -t https://www.baidu.com -f ip.txt -t 1000

// more help info
go run main.go -h
```
