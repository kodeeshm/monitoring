#!/bin/bash

DATE=$(date +'%Y-%m-%d')
LOGNAME=$(date +"%Y-%m-%d-%T" | sed -e 's/:/'_'/g')
LOG_RETENTION="7"

# Write a log file with details
DEVICENAME=`hostname`
TS=$(echo $DATE_$LOGNAME)
FILE=`echo $DEVICENAME-$TS.log`

echo -e "***** HOST INFORMATION *****" >> $FILE
hostnamectl >> $FILE
echo "" >> $FILE
echo "-------------------------------" >> $FILE
# -Dropped packets details
echo -e "***** Dropped Packets *****" >> $FILE
netstat -s |grep -i dropped >> $FILE
echo "" >> $FILE
echo "-------------------------------" >> $FILE
# -tx details:
echo -e "***** tx_packets and rx_packets *****" >> $FILE
ip -s link >> $FILE
echo "" >> $FILE
echo "-------------------------------" >> $FILE
# -System uptime:
echo -e "***** SYSTEM UPTIME *****" >> $FILE
uptime >> $FILE
echo "" >> $FILE
echo "-------------------------------" >> $FILE
echo -e "Done..." >> $FILE

# Purge based on Log Retention duration
find . -maxdepth 1 -type f -ctime -$LOG_RETENTION -exec rm \{\} \;
