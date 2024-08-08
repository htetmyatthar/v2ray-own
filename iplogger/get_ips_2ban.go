package main

import (
	"bufio"
	"errors"
	"log"
	"os"

	. "github.com/v2fly/v2ray-core/v5/iplogger/logger"
)

const (
	activeIPsFilePath = "active_ips.txt"
	IPsToBanFile      = "ips_2ban.txt"
)

func main() {
	if err := logIPsToBan(IPStore, IPsToBanFile); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func logIPsToBan(store map[string]string, fileToLog string) error {
	activeIPs, err := loadActiveIPs()
	if err != nil {
		return errors.New("Unable to load the active IPs log file.")
	}

	// go through the store
	for ip := range store {
		if _, ok := activeIPs[ip]; ok {
			delete(activeIPs, ip)
		}
	}
	// caution: can't use else in the for loop. the ipaddress of connections in IPStore are the latest.
	// solution: create another for loop to log all the ip.
	file, err := os.OpenFile(fileToLog, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("Error unable to load a file to write the IP address to ban.")
	}
	defer file.Close()

	// get all the ips that are not deleted.
	for ip := range activeIPs {
		if _, err := file.WriteString(ip + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// loadActiveIPs load the currently established active ip addresses from a file that is being logged by a `cron` job.
func loadActiveIPs() (map[string]struct{}, error) {
	ips := make(map[string]struct{}, ConnectionCount)
	file, err := os.Open(activeIPsFilePath)
	if err != nil {
		return nil, errors.New("Error loading active tcp connection `IP` addresses file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ips[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error reading active IP addresses file")
	}

	return ips, nil
}
