package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shirou/gopsutil/disk"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
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
		if !stringInList(device, mountPoints) {
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

type TransIn3 struct {
	PayEngineId   int    `db:"pay_engine"`
	PayEngineName string `db:"name"`
	TrnCount      int    `db:"cnt"`
}

type SBMRow57 struct {
	Slave_IO_State                string `db:"Slave_IO_State"`
	Master_Host                   string `db:"Master_Host"`
	Master_User                   string `db:"Master_User"`
	Master_Port                   string `db:"Master_Port"`
	Connect_Retry                 string `db:"Connect_Retry"`
	Master_Log_File               string `db:"Master_Log_File"`
	Read_Master_Log_Pos           string `db:"Read_Master_Log_Pos"`
	Relay_Log_File                string `db:"Relay_Log_File"`
	Relay_Log_Pos                 string `db:"Relay_Log_Pos"`
	Relay_Master_Log_File         string `db:"Relay_Master_Log_File"`
	Slave_IO_Running              string `db:"Slave_IO_Running"`
	Slave_SQL_Running             string `db:"Slave_SQL_Running"`
	Replicate_Do_DB               string `db:"Replicate_Do_DB"`
	Replicate_Ignore_DB           string `db:"Replicate_Ignore_DB"`
	Replicate_Do_Table            string `db:"Replicate_Do_Table"`
	Replicate_Ignore_Table        string `db:"Replicate_Ignore_Table"`
	Replicate_Wild_Do_Table       string `db:"Replicate_Wild_Do_Table"`
	Replicate_Wild_Ignore_Table   string `db:"Replicate_Wild_Ignore_Table"`
	Last_Errno                    string `db:"Last_Errno"`
	Last_Error                    string `db:"Last_Error"`
	Skip_Counter                  string `db:"Skip_Counter"`
	Exec_Master_Log_Pos           string `db:"Exec_Master_Log_Pos"`
	Relay_Log_Space               string `db:"Relay_Log_Space"`
	Until_Condition               string `db:"Until_Condition"`
	Until_Log_File                string `db:"Until_Log_File"`
	Until_Log_Pos                 string `db:"Until_Log_Pos"`
	Master_SSL_Allowed            string `db:"Master_SSL_Allowed"`
	Master_SSL_CA_File            string `db:"Master_SSL_CA_File"`
	Master_SSL_CA_Path            string `db:"Master_SSL_CA_Path"`
	Master_SSL_Cert               string `db:"Master_SSL_Cert"`
	Master_SSL_Cipher             string `db:"Master_SSL_Cipher"`
	Master_SSL_Key                string `db:"Master_SSL_Key"`
	Seconds_Behind_Master         string `db:"Seconds_Behind_Master"`
	Master_SSL_Verify_Server_Cert string `db:"Master_SSL_Verify_Server_Cert"`
	Last_IO_Errno                 string `db:"Last_IO_Errno"`
	Last_IO_Error                 string `db:"Last_IO_Error"`
	Last_SQL_Errno                string `db:"Last_SQL_Errno"`
	Last_SQL_Error                string `db:"Last_SQL_Error"`
	Replicate_Ignore_Server_Ids   string `db:"Replicate_Ignore_Server_Ids"`
	Master_Server_Id              string `db:"Master_Server_Id"`
	Master_UUID                   string `db:"Master_UUID"`
	Master_Info_File              string `db:"Master_Info_File"`
	SQL_Delay                     string `db:"SQL_Delay"`
	SQL_Remaining_Delay           string `db:"SQL_Remaining_Delay"`
	Slave_SQL_Running_State       string `db:"Slave_SQL_Running_State"`
	Master_Retry_Count            string `db:"Master_Retry_Count"`
	Master_Bind                   string `db:"Master_Bind"`
	Last_IO_Error_Timestamp       string `db:"Last_IO_Error_Timestamp"`
	Last_SQL_Error_Timestamp      string `db:"Last_SQL_Error_Timestamp"`
	Master_SSL_Crl                string `db:"Master_SSL_Crl"`
	Master_SSL_Crlpath            string `db:"Master_SSL_Crlpath"`
	Retrieved_Gtid_Set            string `db:"Retrieved_Gtid_Set"`
	Executed_Gtid_Set             string `db:"Executed_Gtid_Set"`
	Auto_Position                 string `db:"Auto_Position"`
	Replicate_Rewrite_DB          string `db:"Replicate_Rewrite_DB"`
	Channel_Name                  string `db:"Channel_Name"`
	Master_TLS_Version            string `db:"Master_TLS_Version"`
}

type SBMRow struct {
	Slave_IO_State                string `db:"Slave_IO_State"`
	Master_Host                   string `db:"Master_Host"`
	Master_User                   string `db:"Master_User"`
	Master_Port                   string `db:"Master_Port"`
	Connect_Retry                 string `db:"Connect_Retry"`
	Master_Log_File               string `db:"Master_Log_File"`
	Read_Master_Log_Pos           string `db:"Read_Master_Log_Pos"`
	Relay_Log_File                string `db:"Relay_Log_File"`
	Relay_Log_Pos                 string `db:"Relay_Log_Pos"`
	Relay_Master_Log_File         string `db:"Relay_Master_Log_File"`
	Slave_IO_Running              string `db:"Slave_IO_Running"`
	Slave_SQL_Running             string `db:"Slave_SQL_Running"`
	Replicate_Do_DB               string `db:"Replicate_Do_DB"`
	Replicate_Ignore_DB           string `db:"Replicate_Ignore_DB"`
	Replicate_Do_Table            string `db:"Replicate_Do_Table"`
	Replicate_Ignore_Table        string `db:"Replicate_Ignore_Table"`
	Replicate_Wild_Do_Table       string `db:"Replicate_Wild_Do_Table"`
	Replicate_Wild_Ignore_Table   string `db:"Replicate_Wild_Ignore_Table"`
	Last_Errno                    string `db:"Last_Errno"`
	Last_Error                    string `db:"Last_Error"`
	Skip_Counter                  string `db:"Skip_Counter"`
	Exec_Master_Log_Pos           string `db:"Exec_Master_Log_Pos"`
	Relay_Log_Space               string `db:"Relay_Log_Space"`
	Until_Condition               string `db:"Until_Condition"`
	Until_Log_File                string `db:"Until_Log_File"`
	Until_Log_Pos                 string `db:"Until_Log_Pos"`
	Master_SSL_Allowed            string `db:"Master_SSL_Allowed"`
	Master_SSL_CA_File            string `db:"Master_SSL_CA_File"`
	Master_SSL_CA_Path            string `db:"Master_SSL_CA_Path"`
	Master_SSL_Cert               string `db:"Master_SSL_Cert"`
	Master_SSL_Cipher             string `db:"Master_SSL_Cipher"`
	Master_SSL_Key                string `db:"Master_SSL_Key"`
	Seconds_Behind_Master         string `db:"Seconds_Behind_Master"`
	Master_SSL_Verify_Server_Cert string `db:"Master_SSL_Verify_Server_Cert"`
	Last_IO_Errno                 string `db:"Last_IO_Errno"`
	Last_IO_Error                 string `db:"Last_IO_Error"`
	Last_SQL_Errno                string `db:"Last_SQL_Errno"`
	Last_SQL_Error                string `db:"Last_SQL_Error"`
	Replicate_Ignore_Server_Ids   string `db:"Replicate_Ignore_Server_Ids"`
	Master_Server_Id              string `db:"Master_Server_Id"`
	Master_UUID                   string `db:"Master_UUID"`
	Master_Info_File              string `db:"Master_Info_File"`
	SQL_Delay                     string `db:"SQL_Delay"`
	SQL_Remaining_Delay           string `db:"SQL_Remaining_Delay"`
	Slave_SQL_Running_State       string `db:"Slave_SQL_Running_State"`
	Master_Retry_Count            string `db:"Master_Retry_Count"`
	Master_Bind                   string `db:"Master_Bind"`
	Last_IO_Error_Timestamp       string `db:"Last_IO_Error_Timestamp"`
	Last_SQL_Error_Timestamp      string `db:"Last_SQL_Error_Timestamp"`
	Master_SSL_Crl                string `db:"Master_SSL_Crl"`
	Master_SSL_Crlpath            string `db:"Master_SSL_Crlpath"`
	Retrieved_Gtid_Set            string `db:"Retrieved_Gtid_Set"`
	Executed_Gtid_Set             string `db:"Executed_Gtid_Set"`
	Auto_Position                 string `db:"Auto_Position"`
}

func checkSBM(host, port, user, pass, dbname string, threshold int) []string {
	var retList []string
	db, err := sqlx.Connect("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("MySQL version: ", version)
	if version[0:3] == "5.7" {
		var sbmRow SBMRow57
		db.QueryRowx("show slave status").StructScan(&sbmRow)
		if sbmRow.Slave_SQL_Running == "No" || sbmRow.Slave_IO_Running == "No" {
			retList = append(retList, "Slave_IO_Running: "+sbmRow.Slave_IO_Running+"\n"+
				"Slave_SQL_Running: "+sbmRow.Slave_SQL_Running+"\n"+
				"Seconds_Behind_Master: "+sbmRow.Seconds_Behind_Master+"\n")
			return retList
		}
		intSBM, err := strconv.Atoi(sbmRow.Seconds_Behind_Master)
		if err != nil {
			panic(err)
		}
		if intSBM > threshold {
			retList = append(retList, "Slave_IO_Running: "+sbmRow.Slave_IO_Running+"\n"+
				"Slave_SQL_Running: "+sbmRow.Slave_SQL_Running+"\n"+
				"Seconds_Behind_Master: "+sbmRow.Seconds_Behind_Master+"\n")
		}
		return retList
	}
	if version[0:3] == "5.6" {
		var sbmRow SBMRow
		db.QueryRowx("show slave status").StructScan(&sbmRow)
		if sbmRow.Slave_SQL_Running == "No" || sbmRow.Slave_IO_Running == "No" {
			retList = append(retList, "Slave_IO_Running: "+sbmRow.Slave_IO_Running+"\n"+
				"Slave_SQL_Running: "+sbmRow.Slave_SQL_Running+"\n"+
				"Seconds_Behind_Master: "+sbmRow.Seconds_Behind_Master+"\n")
			return retList
		}
		intSBM, err := strconv.Atoi(sbmRow.Seconds_Behind_Master)
		if err != nil {
			panic(err)
		}
		if intSBM > threshold {
			retList = append(retList, "Slave_IO_Running: "+sbmRow.Slave_IO_Running+"\n"+
				"Slave_SQL_Running: "+sbmRow.Slave_SQL_Running+"\n"+
				"Seconds_Behind_Master: "+sbmRow.Seconds_Behind_Master+"\n")
		}
		return retList
	}
	return []string{"Unsupported MySQL version"}
}

func checkPayEngine(host, port, user, pass, dbname string, threshold int) []string {
	var retList []string
	db, err := sqlx.Connect("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	sql := `
select p.pay_engine,
       pe.name,
       count(p.id) as cnt
from pays_main as p
left join pays_engines as pe on p.pay_engine=pe.id
where gstatus = 3
group by pay_engine`

	var peCount []TransIn3
	err = db.Select(&peCount, sql)
	if err != nil {
		panic(err)
	}
	for _, row := range peCount {
		if row.TrnCount > threshold {
			retList = append(retList, "PayEngine: "+strconv.Itoa(row.PayEngineId)+"\n"+
				"Name: "+row.PayEngineName+"\n"+
				"Count: "+strconv.Itoa(row.TrnCount)+"\n")
		}
	}
	return retList
}

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.DisableColors = true
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	//log.SetReportCaller(true)
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/j_notif/") // path to look for the config file in
	viper.AddConfigPath("$HOME/.j_notif")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorln("Config not found...")
		log.Errorln(err)
		os.Exit(1)
	}

	tgToken := viper.GetString("tgToken")
	if tgToken == "" {
		log.Errorln("tgToken is required")
		os.Exit(1)
	}
	chatId := viper.GetString("chatId")
	if chatId == "" {
		log.Errorln("chatId is required")
		os.Exit(1)
	}

	disksToCheck := viper.GetStringSlice("disksToCheck")
	if len(disksToCheck) == 0 {
		log.Infoln("disksToCheck is empty")
	} else {
		prevDiskNotif := viper.GetStringSlice("diskNotif")
		diskNotif := diskUsage(disksToCheck, viper.GetFloat64("diskThreshold"))
		if testEq(prevDiskNotif, diskNotif) {
			log.Infoln("same disk notification. do not send")
		} else {
			if len(diskNotif) > 0 {
				text := "disksToCheck\n"
				for _, item := range diskNotif {
					text = text + item + "\n"
				}
				sendMsg(tgToken, chatId, text)
			} else {
				sendMsg(tgToken, chatId, "disksToCheck: OK")
			}
		}
		viper.Set("diskNotif", diskNotif)
	}

	if viper.GetBool("checkSecondsBehindMaster") {
		sbmNotif := checkSBM(viper.GetString("mysqlhost"), viper.GetString("mysqlport"),
			viper.GetString("mysqluser"), viper.GetString("mysqlpass"),
			viper.GetString("mysqldb"), viper.GetInt("mysqlsbmthreshold"))
		prevSBMNotif := viper.GetStringSlice("prevSBMNotif")
		if testEq(prevSBMNotif, sbmNotif) {
			log.Infoln("same SBM notification. do not send")
		} else {
			text := "checkSecondsBehindMaster\n"
			if len(sbmNotif) > 0 {
				for _, item := range sbmNotif {
					text = text + item + "\n"
				}
			} else {
				text = text + "OK\n"
			}
			sendMsg(tgToken, chatId, text)
		}
		viper.Set("prevSBMNotif", sbmNotif)
	}

	if viper.GetBool("check_pay_engine") {
		peNotif := checkPayEngine(viper.GetString("mysqlhost"), viper.GetString("mysqlport"),
			viper.GetString("mysqluser"), viper.GetString("mysqlpass"),
			viper.GetString("mysqldb"), viper.GetInt("pay_engine_threshold"))
		prevPENotif := viper.GetStringSlice("prevPENotif")
		if testEq(prevPENotif, peNotif) {
			log.Infoln("same pay_engine notification. do not send")
		} else {
			text := "check_pay_engine\n"
			if len(peNotif) > 0 {
				for _, item := range peNotif {
					text = text + item + "\n"
				}
			} else {
				text = text + " OK\n"
			}
			sendMsg(tgToken, chatId, text)
		}
		viper.Set("prevPENotif", peNotif)
	}

	viper.WriteConfig()
}
