# SubAnser: Tool to check the website's avilability on various ports.

## Overview

SubAnser is a Go program designed to check the availability of web services across multiple domains and ports. It efficiently handles concurrent requests to quickly determine the status of specified services, making it a valuable tool for administrators and web developers looking to monitor the health of their online platforms.

## Features

- **Concurrent Requests:** Speed up the monitoring process by checking multiple services at the same time.
- **Multiple Domain Support:** Easily check the availability of services across different domains.
- **Customizable Port Checks:** Specify which ports to check for each domain.
- **Timeout Settings:** Modify the timeout settings to suit different network speeds.
- **Error Logging:** Quickly identify and resolve issues with unreachable services.

## Requirements

- Go (Version 1.15 or higher recommended)
- An active internet connection

## Note
This tool is written for my learning purpose as previous script was created with the help of friend. I have tried recreating it in my fav lang GO. It might have errors or might give false positie if you  feel their is some wrong
wil recommnded it checking manually. You can change the protocols by editting manually  as one can also change the inerval set in the reqHandler function which is currently set to 30 sec. 

## Installation

To set up SubAnswer on your machine, follow these steps:

1. Clone the repository or download the source code and navigate to that directory:
   ```bash
   git clone https://yourrepository.com/subanser.git
   cd subanser
2. For the usage
   ```bash
   go run main.go -f domains -p 80,443,8080
It will start making request on every domain on the given port with both protocols HTTP and HTTPS. Error's faced on particular port or domain name will be logged to a diffrent file.

example:
![image](https://github.com/clevercoder91/Subanser/assets/71368140/ed99548d-900b-46f8-82e2-9c75e6d13a47)



