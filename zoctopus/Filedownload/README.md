## 实现思路  

文件下载  

### 想到的问题和细节  

- 1，如何控制下载速度？  
buffer  
[ratelimit包](https://github.com/juju/ratelimit)  

- 2，如何在下载完文件后退出程序？  
设置状态，如果速度为0，说明下载完了，break退出；  
读取文件的`EOF`值。  

### 可扩展性  

- 新建config配置文件，把下载速度，文件网址，版本号做成可以配置的  

#### 运行样例  
```
$ ./Filedownload 
2019/07/22 23:17:46 [*]Filenamego1.10.3.darwin-amd64.pkg
[*] Speed 3.047 MB / 1s 
[*] Speed 6.535 MB / 1s 
[*] Speed 5.919 MB / 1s 
[*] Speed 6.000 MB / 1s 
[*] Speed 5.695 MB / 1s 
[*] Speed 5.715 MB / 1s 
[*] Speed 5.859 MB / 1s 
[*] Speed 5.747 MB / 1s 
[*] Speed 5.972 MB / 1s 
[*] Speed 5.812 MB / 1s 
[*] Speed 4.153 MB / 1s 
[*] Speed 5.941 MB / 1s 
[*] Speed 5.582 MB / 1s 
[*] Speed 5.993 MB / 1s 
[*] Speed 5.972 MB / 1s 
[*] Speed 6.032 MB / 1s 
[*] Speed 5.890 MB / 1s 
[*] Speed 5.969 MB / 1s 
[*] Speed 6.000 MB / 1s 
[*] Speed 5.669 MB / 1s 
[*] Speed 3.831 MB / 1s 
[*] Speed 5.637 MB / 1s 
[*] Speed 554.606 KB / 1s 
[*] Speed 0 Bytes / 1s 
```

## 收获  
- 文件操作  
- http请求  
- 协程  
- 定时器机制

## 参考资料  
[Golang 使用http Client下载文件](https://blog.csdn.net/a99361481/article/details/81751231)