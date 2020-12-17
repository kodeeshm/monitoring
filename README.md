## Monitoring Script for packet dropped, Uptime, RX and TX with Purge enabled
### Place the script the on the server and run the following commands to start...

### For Go
###### go run monitoring.go

### Shell
###### You have to update the common drive details in the script before beginning the steps.
Line no 34: scp $FILE USERNAME@REMOTEHOST:/<LOG_LOCATION>

- **Step1:** 
###### chmod u+x monitoring.sh
- **Step2:** 
###### ./monitoring.sh
- **Step3**
###### sudo vi /etc/crontab
- Copy and Paste the following
###### 5 * * * * root <script_location>
