## wakeup-go

这个项目是用go语言开发，为了能在路由器或其他长开机的兼容设备上运行，以方便在支持WOL的机器上的远程开机。由go-wol代码修改而来，感谢原作者

### 远程开机前需要做的

* 确认设备的cpu型号
   
   一般来说可以用如下命令查看cpu类型
   ```
   cat /proc/cpuinfo | grep 'model' | awk -F ":" '{print $2}'
   ```
   再下载对应二进制文件

* 开启BIOS中WOL的功能
  
  不同的BIOS开启的方式不同，请自行搜索对应方法

* 最好关闭需远程开机的机器的睡眠功能

### 如何运行

* 打开wakesh.sh文件，第二行中在MAC_ADDR=后填入要开启机器的mac地址
* 如果是在路由上运行，请root路由，确保ssh可登陆，并拷贝对应的二进制文件及wakesh.sh
* 运行命令`chmod +x wakesh.sh chmod +x wakeup ./wakesh.sh`

### 对应的二进制文件下载

  ![mips wakeup](https://github.com/oserz/wakeup-go/tree/master/bin/mips/wakeup)

  ![mipsle wakeup](https://github.com/oserz/wakeup-go/tree/master/bin/mipsle/wakeup)

  ![arm wakeup](https://github.com/oserz/wakeup-go/tree/master/bin/arm/wakeup)

  ![x86 wakeup](https://github.com/oserz/wakeup-go/tree/master/bin/x86/wakeup)

  ![amd64 wakeup](https://github.com/oserz/wakeup-go/tree/master/bin/amd64/wakeup)


### 最后

如果这个项目对您有帮助，请点个赞，当然如果能打赏一点就更感谢啦

![](/img/zfb.jpg)

![](/img/wx.jpg)