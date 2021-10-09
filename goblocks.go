package main

// replace  goblocks main file
import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"io/ioutil"
	"log"

	"strconv"
	"strings"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/pelletier/go-toml"
)

const (
	iconCPU      = " "
	iconRAM      = " "
	iconUp       = " "
	iconDown     = "  "
	iconBri      = " ﯦ "
	iconKeyboard = ""
	iconPlug     = " "
	iconVpn      = " "
)

var (
	iconTimeArr = [12]string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
	iconBatArr  = [5]string{" ", " ", " ", " ", " "}
	iconVolArr  = [4]string{"", "", "墳", ""}
	netDevMap   = map[string]struct{}{}
	cpuOld, _   = cpu.Get()
	rxOld       = 0
	txOld       = 0
	wlan        = "wlan0"
	lan         = "lo"
	style       = "background"
	// netColor    = "#d08070"
	// cpuColor    = "#ebcb8b"
	// memColor    = "#a3be8c"
	// volColor    = "#5e81ac"
	// batColor    = "#88c0d0"
	// datColor    = "#b48ead"
)

func main() {
	// parseConfig()
	for {
		status := setStyle(style)
		s := strings.Join(status, " ")
		exec.Command("xsetroot", "-name", s).Run()

		var now = time.Now()
		time.Sleep(now.Truncate(time.Second).Add(time.Second).Sub(now))
	}
}

func setStyle(style string) []string {
	// var briefStyle string
	// if style == "background" {
	// 	briefStyle = "^b"
	// } else {
	// 	briefStyle = "^c"
	// }

	return []string{
		// briefStyle + netColor + "^",
		// updateNet(),
		updateBri(),
		updateKey(),
		// briefStyle + cpuColor + "^",
		updateCPU(),
		// briefStyle + memColor + "^",
		updateMem(),
		// briefStyle + volColor + "^",
		updateVolume(),
		// briefStyle + batColor + "^",
		updateBattery(),
		// briefStyle + datColor + "^",
		updateDateTime(),
		updateVpn(),
	}
}
func updateVpn() string {
	cmd := exec.Command("systemctl", "check", "clash.service")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
		} else {
			fmt.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
		}
	}
	fmt.Printf(string(out))
	res := strings.Trim(string(out), "\n")
	if res == "active" {
		return iconVpn
	} else {
		return ""
	}
}
func updateBri() string {
	var bri []byte
	var maxBri []byte
	var err error
	var getCurrentBri *exec.Cmd
	var getMaxBri *exec.Cmd

	// 执行单个shell命令时, 直接运行即可
	getCurrentBri = exec.Command("cat", "/sys/class/backlight/intel_backlight/brightness")
	if bri, err = getCurrentBri.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bri_f, err := strconv.Atoi(strings.Trim(string(bri), "\n"))

	getMaxBri = exec.Command("cat", "/sys/class/backlight/intel_backlight/max_brightness")
	if maxBri, err = getMaxBri.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	maxBri_f, err := strconv.Atoi(strings.Trim(string(maxBri), "\n"))
	per := int(float32(bri_f) / float32(maxBri_f) * 100)
	// 默认输出有一个换行
	// fmt.Println(string(whoami))
	// 指定参数后过滤换行符
	fmt.Println(strings.Trim(string(bri), "\n"))
	return iconBri + strconv.Itoa(per) + "%"
}

func updateKey() string {
	var layout []byte
	var err error
	var cmd *exec.Cmd

	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("bash", "-c", "setxkbmap -print | grep xkb_symbols | sed -r 's/.*\\+(..)\\+.*/\\1/'")
	if layout, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(strings.Trim(string(layout), "\n"))
	res := strings.Trim(string(layout), "\n")
	return iconKeyboard + " " + strings.ToUpper(res)
}

func getNetSpeed() (int, int) {
	dev, err := os.Open("/proc/net/dev")
	if err != nil {
		log.Fatalln(err)
	}
	defer dev.Close()

	devName, rx, tx, rxNow, txNow, void := "", 0, 0, 0, 0, 0
	for scanner := bufio.NewScanner(dev); scanner.Scan(); {
		_, err = fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%s %d %d %d %d %d %d %d %d %d", &devName, &rx, &void, &void, &void, &void, &void, &void, &void, &tx)
		if _, ok := netDevMap[devName]; ok {
			rxNow += rx
			txNow += tx
		}
	}
	fmt.Println(rxNow, txNow)
	return rxNow, txNow
}

func updateNet() string {
	rxNow, txNow := getNetSpeed()
	defer func() { rxOld, txOld = rxNow, txNow }()
	return iconDown + fmtNetSpeed(float64(rxNow-rxOld)) + " " + iconUp + fmtNetSpeed(float64(txNow-txOld))

}

func fmtNetSpeed(speed float64) string {
	if speed < 0 {
		log.Fatalln("Speed must be positive")
	}
	var res string

	switch {
	case speed >= (1024 * 1024 * 1024):
		gbSpeed := speed / (1024.0 * 1024.0 * 1024.0)
		res = fmt.Sprintf("%.2f", gbSpeed) + "Gb"
	case speed >= (1024 * 1024):
		mbSpeed := speed / (1024.0 * 1024.0)
		res = fmt.Sprintf("%.1f", mbSpeed) + "Mb"
	case speed >= 1024:
		kbSpeed := speed / 1024.0
		res = fmt.Sprintf("%.1f", kbSpeed) + "kb"
	case speed >= 0:
		res = fmt.Sprint(speed) + "B"
	}

	return res
}

func updateMem() string {
	meminfo, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatalln(err)
	}
	defer meminfo.Close()

	// var total, avail float64
	// for info := bufio.NewScanner(meminfo); info.Scan(); {
	// 	key, value := "", 0.0
	// 	if _, err = fmt.Sscanf(info.Text(), "%s %f", &key, &value); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	if key == "MemTotal:" {
	// 		total = value
	// 	}
	// 	if key == "MemAvailable:" {
	// 		avail = value
	// 	}
	// }
	var avail float64
	for info := bufio.NewScanner(meminfo); info.Scan(); {
		key, value := "", 0.0
		if _, err = fmt.Sscanf(info.Text(), "%s %f", &key, &value); err != nil {
			log.Fatalln(err)
		}
		if key == "MemAvailable:" {
			avail = value
		}
	}
	// used := (total - avail) / 1024.0 / 1024.0
	avail = avail / 1024.0 / 1024.0

	return iconRAM + fmt.Sprintf("%.2f", avail) + "GiB"
}

func updateCPU() string {
	cpuNow, err := cpu.Get()
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer func() { cpuOld = cpuNow }()
	total := float64(cpuNow.Total - cpuOld.Total)
	usage := 100.0 - float64(cpuNow.Idle-cpuOld.Idle)/total*100
	return iconCPU + fmt.Sprintf("%.2f", usage) + "%"
}

func updateVolume() string {
	const pamixer = "pamixer"
	isMuted, _ := strconv.ParseBool(cmdReturn(pamixer, "--get-mute"))
	volume := cmdReturn(pamixer, "--get-volume")
	if isMuted {
		return iconVolArr[0]
	} else {
		return getVolIcon(volume) + " " + volume + "%"
	}
}

func getVolIcon(volume string) string {
	var res string
	volumeInt, _ := strconv.ParseInt(volume, 10, 32)
	if volumeInt > 80 {
		res = iconVolArr[3]
	} else if volumeInt > 50 {
		res = iconVolArr[2]
	} else if volumeInt > 20 {
		res = iconVolArr[1]
	} else {
		res = iconVolArr[0]
	}
	return res
}

func cmdReturn(bin string, arg string) string {
	var res string
	cmd := exec.Command(bin, arg)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	res = strings.TrimSpace(string(stdout.Bytes()))

	return res
}

func updateBattery() string {
	const pathToPowerSupply = "/sys/class/power_supply/"
	// var pathToBat0 = pathToPowerSupply + "BAT0/"
	// var pathToAC = pathToPowerSupply + "AC/"
	var pathToBat0 = pathToPowerSupply + "BAT0/"
	var pathToAC = pathToPowerSupply + "ADP1/"

	status := parseTxt(pathToBat0, "status")
	capacity := parseTxt(pathToBat0, "capacity")
	isPlugged, _ := strconv.ParseBool(parseTxt(pathToAC, "online"))
	if status == "Full" {
		return iconBatArr[4] + " Full"
	} else {
		if isPlugged == true {
			// return getBatIcon(capacity) + " " + capacity + "%"
			return iconPlug + " " + capacity + "%"
		} else {
			return getBatIcon(capacity) + " " + capacity + "%"
		}
	}
}

func getBatIcon(capacity string) string {
	var res string
	capacityInt, _ := strconv.ParseInt(capacity, 10, 32)
	if capacityInt >= 75 {
		res = iconBatArr[3]
	} else if capacityInt > 50 {
		res = iconBatArr[2]
	} else if capacityInt > 25 {
		res = iconBatArr[1]
	} else {
		res = iconBatArr[0]
	}
	return res
}

func parseTxt(path string, name string) string {
	var res string
	contentOri, err := ioutil.ReadFile(path + name)
	if err != nil {
		log.Println("Please check the " + name + "'s path")
	}
	res = strings.TrimSpace(string(contentOri))

	return res
}

func parseConfig() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	path := dirname + "/.config/goblocks/config.toml"

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	config, _ := toml.Load(string(content))
	wlan = config.Get("networks.wlan").(string) + ":"
	lan = config.Get("networks.lan").(string) + ":"

	style = config.Get("color.style").(string)
	// netColor = config.Get("color.netColor").(string)
	// cpuColor = config.Get("color.cpuColor").(string)
	// memColor = config.Get("color.memColor").(string)
	// volColor = config.Get("color.volColor").(string)
	// batColor = config.Get("color.batColor").(string)
	// datColor = config.Get("color.datColor").(string)

	netDevMap[wlan] = struct{}{}
	netDevMap[lan] = struct{}{}
}

func updateDateTime() string {
	var hour = time.Now().Hour()
	var dateTime = time.Now().Local().Format("2006-01-02 Mon 15:04:05")

	return getHourIcon(hour) + dateTime
}

func getHourIcon(hour int) string {
	var res string
	if hour == 0 || hour == 12 {
		res = iconTimeArr[11]
	} else if hour == 23 || hour == 11 {
		res = iconTimeArr[10]
	} else if hour == 22 || hour == 10 {
		res = iconTimeArr[9]
	} else if hour == 21 || hour == 9 {
		res = iconTimeArr[8]
	} else if hour == 20 || hour == 8 {
		res = iconTimeArr[7]
	} else if hour == 19 || hour == 7 {
		res = iconTimeArr[6]
	} else if hour == 18 || hour == 6 {
		res = iconTimeArr[5]
	} else if hour == 17 || hour == 5 {
		res = iconTimeArr[4]
	} else if hour == 16 || hour == 4 {
		res = iconTimeArr[3]
	} else if hour == 15 || hour == 3 {
		res = iconTimeArr[2]
	} else if hour == 14 || hour == 2 {
		res = iconTimeArr[1]
	} else if hour == 13 || hour == 1 {
		res = iconTimeArr[0]
	}
	return res
}
