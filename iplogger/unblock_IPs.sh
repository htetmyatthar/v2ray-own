#!/bin/bash

# File containing the list of IP addresses to unblock
BLOCKED_IPS_FILE="blocked_ips.txt"
PORT="80"  # Specify the port used in your rules

# Function to remove IP blocking rules for a specific port
unblock_ips() {
  while IFS= read -r IP; do
	if [[ -n "$IP" ]]; then
		sudo iptables -D INPUT -p tcp -s "$IP" --dport "$PORT" -j DROP
	fi 
  done < "$BLOCKED_IPS_FILE"

  # clear the log file after unblocking
  : > "$BLOCKED_IPS_FILE"
}

# Unblock the IPs
unblock_ips
