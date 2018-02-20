# sni-scaner
Fastest Server Name Indication(SNI) proxy scaner

# Usage
```
echo '1.1.1.1/24' > cn.txt
go run sni.go -h
go run sni.go -t https://www.baidu.com -f cn.txt -d 1500 > res.log
```
