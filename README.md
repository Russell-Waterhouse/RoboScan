# RoboScan

## About the project
Inspired by GoBuster, RoboScan is a CLI website enumeration tool that uses a 
website's robots.txt as the wordlist.   
It makes a request to every listing in the /robots.txt file with the "Disallow" Tag.   


Because the robots.txt file is often a list of everything that website's administrator
doesn't want the public internet to know about, a tool that quickly makes 
requests to every listing in robots.txt and tells you which are accessible is a handy little thing to have. 
This tool is useful to red teams doing security assessments on websites. 

## Using the Project
To execute the program in a Linux environment, simply run the compiled executable in this GitHub repository. 
```Bash
wget https://github.com/Russell-Waterhouse/RoboScan/raw/release1.0/roboscan1.0
chmod +x roboscan1.0
./roboscan1.0
```

You'll be met with the prompt: `What Website do you want to scan?`

Simply type in your response, for example `example.com`, and the script will retrieve the file at `https://example.com/robots.txt` and 
print the results of requesting every "Disallow" entry.   
  
  Happy Hacking!
  
## A Concluding Note: 
My employers and I take no responsibility, personally or comercially, for malicious use of this tool. It is intended to be used by security assessment
professionals and website maintainers to assess the security posture of websites that they own and maintain. Do not use this tool on websites that you 
do not have permission to test on. 

