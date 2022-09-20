#!/usr/bin/env python3
import requests as request
import json
# test01
# route: /post/update/{id}
# method: put
# expected result: returns 200





host = "127.0.0.1"
port = "8080"



def update_post_test():
    route = '/post/update/'
    custom_id = input('Input post id: ')
    if(custom_id == ''):
        custom_id = "f0668f75-b99e-447f-8495-fbe660bb1892" 
    
    route += custom_id
    data = {
        "title":"Wubba Lubba Dub Dub",
        "content": "R&M FTW!"
    }

    url = "http://{}:{}{}".format(host,port, route)

    rsp = request.put(url, data=json.dumps(data))
    print(rsp.content)
    if (rsp.status_code != 200):
        return -1
    return 0



def main():
    test01 = update_post_test()
    if(test01 == 0):
        print("test01 has passed")
    else:
        print("test01 has failed")

   


main()
