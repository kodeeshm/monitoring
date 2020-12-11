# Monitoring Script with Purge enabled
## Place the script the on the server and run the following commands to start...

- **Step1:** 
chmod u+x monitoring.sh
- **Step2:** 
./monitoring.sh
- **Step3**
sudo vi /etc/crontab
- Copy and Paste the following
- 5 * * * * root <script_location>
