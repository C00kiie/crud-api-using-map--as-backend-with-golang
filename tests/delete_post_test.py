#!/usr/bin/env python3
import requests as request
import json




# test01
# route: /post/delete/{id}
# method: POST
# expected result: returns 200





host = "127.0.0.1"
port = "8080"



def delete_post_by_id():
    route = '/post/delete/'
    custom_id = input("Input the ID: ")
    route += custom_id

    url = "http://{}:{}{}".format(host,port, route)

    print(route)
    rsp = request.delete(url)
    print(rsp.content)
    if (rsp.status_code != 200):
        return -1
    return 0



def main():
    test01 = delete_post_by_id()
    if(test01 == 0):
        print("test01 has passed")
    else:
        print("test01 has failed")

   


main()
