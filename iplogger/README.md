# iplogger
This is a simple ip logging mechanism to use with **vmess** inbound protocol.

# Architecture
`block_IPs.sh` will log the active tcp connections and block all of the excessive connection.
This will block the IP address that is using the same **alterId** for the connection.

`unblock_IPs.sh` will unblock all the blocked connections.

# Setting up the automation

## Setting up the banner
logger logs the active connections' IP addresses to be compared with the latest **vaild IP addresses** that is logged with alterId. 

1) create executable of `get_ips_2ban`
    ```bash
    go build -o get_ips_2ban /path/to/iplogger/
    ```

2) make sure that your server is using IPv4 or IPv6 and choose a logging command in `log_active_IPs.sh` file.

3) make sure that the script is executable by
    ```bash
    chmod +x /path/to/block_IPs.sh
    ```

4) make the script automatively running using `crontab`
    ```bash
    sudo crontab -e 
    */10****/path/to/block_IPs.sh
    ```

5) make sure that the cron job's been added by
    ```bash
    crontab -l
    ```

## Setting up the IP unblocker
Unblocker remove the banned IP address after some time.

1) make sure that the script is executable by
    ```bash
    chmod +x /path/to/unblock_IPs.sh
    ```

2) make the script automatively running using `crontab`
    ```bash
    sudo crontab -e
    */10****/path/to/unblock_IPs.sh
    ```

3) make sure that the cron job's been added by
    ```bash
    crontab -l
    ```
