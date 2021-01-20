#!/usr/bin/python
# -*- coding: utf-8 -*-
import requests
import time
import optparse

fileObj = open('sample.txt', 'r') #give your file name instead sample.txt
conv_as_arr = fileObj.read().splitlines()

# print(conv_as_arr)

for x in conv_as_arr:
    try:
        response = requests.get('http://' + x)
        print (x + '  ' + str(response))
        time.sleep(1)
    except requests.exceptions.ConnectionError:
        pass
    except requests.exceptions.InvalidURL:
        pass

