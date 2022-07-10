package main

import (
    "bytes"
    "fmt"
    "github.com/shirou/gopsutil/disk"
    "github.com/spf13/viper"
    "net/http"
    "os"
)

func sendMsg(tgToken, chatId, text string) {
    httpposturl := "https://api.telegram.org/bot" + tgToken + "/sendMessage"
    //fmt.Println("HTTP JSON POST URL:", httpposturl)

    hostname, err := os.Hostname()
    if err != nil {
        panic(err)
    }

    textToSend := "Host: " + hostname + "\n" + text

    var jsonData = []byte(`{
		"chat_id": "` + chatId + `",
		"text": "` + textToSend + `"
	}`)
    request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
    request.Header.Set("Content-Type", "application/json; charset=UTF-8")

    client := &http.Client{}
    response, error := client.Do(request)
    if error != nil {
        panic(error)
    }
    defer response.Body.Close()
}

func stringInList(s string, list []string) bool {
    for _, item := range list {
        if s == item {
            return true
        }
    }
    return false
}

func diskUsage(mountPoints []string, threshold float64) []string {
    var retList []string
    parts, _ := disk.Partitions(true)
    for _, p := range parts {
        device := p.Mountpoint
        if false == stringInList(device, mountPoints) {
            continue
        }
        s, _ := disk.Usage(device)

        if s.Total == 0 {
            continue
        }
        if s.UsedPercent > threshold {
            percent := fmt.Sprintf("%2.f%%", s.UsedPercent)
            txt := fmt.Sprintf("Mounted: %s\nUsed: %s\n", p.Mountpoint, percent)
            retList = append(retList, txt)

        }

    }
    return retList
}

func testEq(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func main() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/etc/j_notif/") // path to look for the config file in
    viper.AddConfigPath("$HOME/.j_notif")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("Config not found...")
        os.Exit(1)
    }

    tgToken := viper.GetString("tgToken")
    if tgToken == "" {
        fmt.Println("tgToken is required")
        os.Exit(1)
    }
    chatId := viper.GetString("chatId")
    if chatId == "" {
        fmt.Println("chatId is required")
        os.Exit(1)
    }

    disksToCheck := viper.GetStringSlice("disksToCheck")
    if len(disksToCheck) == 0 {
        fmt.Println("disksToCheck is empty")
    }
    prevDiskNotif := viper.GetStringSlice("diskNotif")
    diskNotif := diskUsage(disksToCheck, 80)
    if testEq(prevDiskNotif, diskNotif) {
        fmt.Println("same disk notification. do not send")
    } else {
        if len(diskNotif) > 0 {
            var text string
            for _, item := range diskNotif {
                text = text + item + "\n"
            }
            sendMsg(tgToken, chatId, text)
        }
    }
    viper.Set("diskNotif", diskNotif)
    viper.WriteConfig()
}
