# ServiceScanner
用golang编写的简单端口探测以及服务探测工具


## 使用方法
帮助
```bash
./serviceScanner -h
```

基本使用
```bash
// 扫描单个ip 
./serviceScanner -h xxx.xx.xx.xx 
// 从文件中读取ip列表
./serviceScanner -hf /home/annevi/ip.txt
```
扫描服务
```bash
./serviceScanner -h xxx.xxx.xxx.xxx -s
```

输出扫描结果
```bash
./serviceScanner -h xxx.xxx.xxx.xxx -o ./res.txt
```

跳过存活探测
```bash
./serviceScanner -h xxx.xxx.xxx.xxx -si 
```

