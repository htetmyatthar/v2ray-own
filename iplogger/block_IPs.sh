#!/bin/bash

# File containing the list of IP addresses
IPS_TO_BAN_FILE="ips_2ban.txt"
ACTIVE_IPS_FILE="active_ips.log"
BLOCKED_IPS_FILE="blocked_ips.log"

PORT="80"
V2RAY_PID=$(pgrep v2ray)

if [ -z "$V2RAY_PID" ]; then
	echo "v2ray process not found. Run the vpn first."
	exit 1
fi


# Ensure the log file exists
touch $ACTIVE_IPS_FILE $BLOCKED_IPS_FILE

# clearing all the data inside the active_ips.log
: > "$ACTIVE_IPS_FILE"

# Get the list of active incoming IP addresses on port 443
# sudo ss -tnp state established -4 | grep "$PORT" | grep "$V2RAY_PID" | awk '{print $4}' | awk -F: '{print $1}' | sort | uniq >> $LOGFILE

# Get the list of active incoming IP addresses on port 443 in ipv6
sudo ss -tnp state established -6 | grep "$PORT" | grep "$V2RAY_PID" | awk '{print $4}' | grep -oP '\[::ffff:\K[^\]]*' | sort | uniq >> $ACTIVE_IPS_FILE

# run the executable to get the ips_2ban.txt
./get_ips_2ban

# Function to block IPs on a specific port using iptables
block_ips() {
  while IFS= read -r IP; do
	if [[ -n "$IP" ]]; then
		sudo iptables -A INPUT -p tcp -s "$IP" --dport "$PORT" -j DROP
		echo "$IP" >> "$BLOCKED_IPS_FILE"
	fi
  done < "$IPS_TO_BAN_FILE"
}

# Block the IPs
block_ips
