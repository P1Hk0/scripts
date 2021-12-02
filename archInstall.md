# archlinux安装(kde)

## 安装arch
https://www.youtube.com/watch?v=pHzROIf0Fuc
https://archlinuxstudio.github.io/ArchLinuxTutorial/#/rookie/basic_install
### 准备
1. rufus烧录
2. 根据电脑机型进入bios关闭secure boot，启动方式uefi，硬盘启动顺序
### 安装
`setfont /usr/share/kbd/consolefonts/LatGrkCyr-12x22.psfu.gz`
1. 确保uefi`ls /sys/firmware/efi/efivars`
2. 连接网络
```
ip link set wlan0 up #比如无线网卡看到叫 wlan0
iwctl                           #进入交互式命令行
device list                     #列出设备名，比如无线网卡看到叫 wlan0
station wlan0 scan              #扫描网络
station wlan0 get-networks      #列出网络 比如想连接CMCC-5AQ7这个无线
station wlan0 connect CMCC-5AQ7 #进行连接 输入密码即可
exit                            #成功后exit退出
dhcpcd
ping baidu.com
```
3. 系统时钟
```
timedatectl set-ntp true    #将系统时间与网络时间进行同步
timedatectl status          #检查服务状态
```
4. 换源
`vim /etc/pacman.d/mirrorlist`

5. 分区
```
fdisk -l 
fdisk /dev/.. # 进入分区了的磁盘里
```

## 安装kde

- install `plasma-meta` dont't install other, too heavy
- [安装教程](https://www.youtube.com/watch?v=BfqbFrE--Bc)

## clash代理()

### 直接执行

```yaml
cd /etc/ && mkdir clash && cd
wget https://github.com/Dreamacro/clash/releases/download/v1.4.1/clash-linux-amd64-v1.4.1.gz
gunzip clash-linux-amd64-v1.4.1.gz
chmod +x clash-linux-amd64-v1.4.1
mv clash-linux-amd64-v1.4.1 /usr/bin/clash
cd /usr/lib/systemd/system
echo "[Unit] 
Description=clash
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
Type=simple
User=root
Group=root
DynamicUser=true
ExecStart=/usr/bin/clash -d /etc/clash/
Restart=always
LimitNOFILE=512000

[Install]
WantedBy=multi-user.target
" >> /usr/lib/systemd/system/clash.service
```

### 以上代码运行完毕后执行

```yaml
cd /usr/bin
./clash点击拷贝拷贝失败拷贝成功
```

然后我们停止运行执行 ctrl+C

首次运行结束后将在用户目录文件下生成 /root/.config/clash 默认配置文件路径，默认配置文件中有config.yaml 这个文件是软件自动生成，我们需要手动添加配置，添加方式如下

随后我们将 Clash 的订阅地址复制到浏览器下载至桌面打开并且复制文件内容。

将下载回来并且复制的 Clash 的配置文件内容全部黏贴至/root/.config/clash/config.yaml 中 文件名不能修改必须是config.yaml

配置文件修改完毕后执行

```yaml
cd /root/.config/clash
cp config.yaml /etc/clash/
cp Country.mmdb /etc/clash/
```

至此重新启动本软件执`sudo systemctl restart clash`

随后系统设置，网络设置中添加http代理，IP 127.0.0.1 端口 7890 socks代理，IP 127.0.0.1 端口 7891

然后通过浏览器访问 [http://clash.razord.top](http://clash.razord.top/) 进行策略组设置

#### 安装后打开方法

1. 执行顺序

```
sudo systemctl start clash
```

2. 系统设置，proxy设置为手动：

   HTTP：127.0.0.1: 7890

   SOCK：127.0.0.1: 7891

3. 打开 http://clash.razord.top/

## 分辨率

直接家目录下创建 ~/.Xresources，写入

```
Xft.dpi:120
```

## 改键

1. 使用Xmodmap （开机启动速度非常慢，5min+）

2. 每次开机手动输入，或者设置自启动文本（失败）

   ```
   setxkbmap -option caps:escape
   ```

## 桌面美化

1. latte-dock

   latte的layout选择plasmapple，需要上kde store下载

2. theme

   系统的global theme选择whitesur，同样需要下载，下载主题可以先下载ocs-url（pacman）

3. virtual desktop

4. kvantum 也是选择whitesur

5. 桌面文本使用krohnkeit (KDE下窗口管理器)

## ZSH设置

1. 安装zsh和oh-my-zsh，主题为默认（网上教程）

2. 插件下载：

   zsh-syntax-highlighting
   zsh-autosuggestions
   
3. 编辑 .zshrc，添加绑定键

   
   
4. 编辑好 后 source .zshrc

##  手机端与电脑数据交互

###  1. 安装resilio sync

```
1. yay -S rslsync
2. mkdir ~/.config/rslsync && touch rslsync.conf
3. rslsync --dump-sample-config > ~/.config/rslsync/rslsync.conf
4. nano rslsync.conf
5. 修改login: admin, password: 123
```

#### 运行

rslsync --config ~/.config/rslsync/rslsync.conf（在zsh设置快捷键 sync）

运行后在浏览器打开 localhost:8888

手机软件扫码传输

#### 问题

用sync只能传送文件，如果要复制文字只能传送邮箱 

### 2. 使用telegram

连接的时候点击左下角的圈圈，选择代理，代理的地址如设置clash的HTTPdigit一样

### 3. onedrive

```
sudo pacman -S onedrive-git
```

[安装后的使用教程](https://www.youtube.com/watch?v=jz5-CtN-WiQ&t=303s) (Youtube)

- 建议只选择一个文件进行同步，因为同步的方式是需要下载到本机。创建一个文件夹linux-data选择同步
- 每次修改玩输入`onedrive`自动同步

## 资源管理器
- htop

## dwm
### 安装
1. 编译之后可使用
2. 需要在登录界面添加启动图标
### dependency
1. alsa-utils
2. acpi-tools
3. pamixer

### status bar

- 可用两种，在`/scripts/autostart.sh`中修改启动设置

1. 使用goblock需要使用go，并且替换`~/goblocks/`中的`goblock.go`文件为`scripts/goglocks.go`
   - 可能需要修改电池方法里的路径（如果电池不显示）
2. 使用`dwm-status.sh`

### 脚本管理
- 在dwm中设置快捷键启动的脚本需要设置启动权限才能启动

    `chmod 755 scripts.sh`

- 设置音量的脚本需要先安装`acpi`并且在`/etc/systemd/system/`中创建一个`addper.service`并添加一下内容

  ```
  [Unit]
  Description=Adding permissions to users
  
  
  [Service]
  ExecStart=/usr/bin/bash -c 'sudo chmod -R 666 /sys/class/backlight/intel_backlight/brightness'
  
  [Install]
  WantedBy=multi-user.target 
  
  ```

  然后启动并查看

  ```shell
  systemctl enable addper.service
  systemctl start addper.service
  systemctl status addper.service
  ```

  目的是开机自动设置目标文件的权限才能修改亮度

## mysql(mariaDb)
### 安装mariadb，几乎与mysql一样
`sudo pacman -S dbmaria`

###　安装服务、初始化
`sudo mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql`

###　加入守护进程。开机自动启动服务
`systemctl enable mysqld`

###　启动服务
`systemctl start mysqld`

### 修改登录密码
`sudo mysqladmin -u root password 123456`

### 登录mysql
`mysql -uroot -p＃　以下为mysql语法，末尾需要加上分号`

### 打开mysql数据
`use mysql;`

### 给用户加权限
`GRANT ALL PRIVILEGES ON *.* TO root@% IDENTIFIED BY 123456;`

### 退出
`exit;`

### 安装dbeaver可视化管理
`sudo pacman -S dbeaver`
