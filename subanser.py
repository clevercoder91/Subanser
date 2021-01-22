#!/usr/bin/python

import requests
import sys, getopt
from ansimarkup import ansiprint as print

def main(argv):
    fileName = ""
    try:
        opts, args = getopt.getopt(argv, "hi:o:")
        if len(opts) == 0:
            print("<red>[-] Usage subanser.py -i <inputfile></red>")
            return
    except getopt.GetoptError:
        print("<red>[-] Missing Required Parameter -i</red>")
        sys.exit(2)
    for opt, arg in opts:
        if opt == "-h":
            print("<red>Usage subanser.py -i <inputfile></red>")
            sys.exit()
        elif opt in ("-i", "--ifile"):
            fileName = arg
    try:
        with open(fileName, "r") as f:
            urls = f.read().splitlines()
            for url in urls:
                try:
                    response = requests.get("http://" + url)
                    print(f"<green>[{url}] - SITE FOUND - <STATUS_CODE> - [{response.status_code}]</green>")    
                except requests.exceptions.ConnectionError:
                    print(f"<red>[{url}] - SITE NOT FOUND - <STATUS_CODE> - [{response.status_code}]</red>")
                except requests.exceptions.InvalidURL:
                    pass
    except FileNotFoundError:
        print(f"<red>[-] No such file named {fileName}</red>")

if __name__ == "__main__":
    main(sys.argv[1:])
