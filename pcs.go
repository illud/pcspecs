package pc

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"golang.org/x/sys/windows"
)

type SysInfo struct {
	Hostname  string `json:"hostname"`
	Platform  string `json:"platform"`
	OsNumber  string `json:"osNumber"`
	CPU       string `json:"cpu"`
	GPU       string `json:"gpu"`
	RAM       string `json:"ram"`
	Disk      int64  `json:"disk"`
	MAINBOARD string `json:"mainboard"`
}

// SysInfo saves the system information
func Spec() SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	// diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	info := new(SysInfo)

	//Extract os number
	platform := hostStat.Platform
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	submatchall := re.FindAllString(platform, -1)
	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.OsNumber = submatchall[0]
	info.CPU = cpuStat[0].ModelName
	info.RAM = strconv.FormatUint(vmStat.Total/1024/1024, 10)[0:2]

	// Gets GPU info
	videoController := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	videoController.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	videoControllerHistory, _ := videoController.Output()
	pcGPU := strings.Replace(string(videoControllerHistory), "Name", "", -1)
	pcGPUString := strings.Replace(pcGPU, "LuminonCore IDDCX Adapter", "", -1)
	info.GPU = pcGPUString

	// Gets MAINBOARD info
	mainBoard := exec.Command("cmd", "/C", "wmic path win32_BaseBoard get Product")
	mainBoard.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	mainBoardHistorys, _ := mainBoard.Output()
	mainBoardString := strings.Replace(string(mainBoardHistorys), "Product", "", -1)
	info.MAINBOARD = mainBoardString

	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64

	_, _, err := c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&freeBytes)))

	if err != nil {
		fmt.Println(err)
	}

	info.Disk = freeBytes / 1e+9

	return *info
}
