package logger

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

// IPLogger is a custom implementation to log a conenction `ip` with
// the coressponding `alterId` so that the we can limit an `ip` to a device.
// therfore we can now treat an `alterId` as one device.

// average number of connection that this v2ray server might make.
const ConnectionCount = 20

// IPStore maps the ID of the client to the *recently* connected IP address.
var IPStore map[string]string = make(map[string]string, ConnectionCount)

// UUIDStore maps the uuids to empty structs for checking that the current UUID is already connected or not.
var UUIDStore map[string]struct{} = make(map[string]struct{}, ConnectionCount)

// LogIP logs the IP of the connection all the time to be used with IP checking.
func LogIP(ID string, ip net.Addr) {
	// only the address
	ipString := strings.Split(ip.String(), ":")[0]
	IPStore[ipString] = ID[:4]

    // Open the file in append mode, create it if it doesn't exist
    file, _ := os.OpenFile("connection.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer file.Close()

    // Write the message to the file
	_, err := file.WriteString(ID + ", " + ipString + "\n")
    if err != nil {
        fmt.Println("error writing to file: %w", err)
    }
	return 
}

// MarkUUID checks the ip of the incoming request connection using the first 4 characters of the uuid.
func MarkUUID(ID string) error {
	hashedUUID := ID[:4]

	// if the uuid is already in use return an error
	if _, ok := UUIDStore[hashedUUID]; ok {
		return errors.New("Given UUID is already in used.")
	}

	UUIDStore[hashedUUID] = struct{}{}
	return nil
}

// UnmarkUUID unmarks the uuid as a connected connection.
func UnmarkUUID(ID string) {
	delete(UUIDStore, ID[:4])
}
