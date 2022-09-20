#!/usr/bin/env python3
import requests as request
import json
# test01
# route: /post/create
# method: POST
# expected result: returns 200





host = "127.0.0.1"
port = "8080"



def create_post_test():
    route = '/post/create'
    data = {
        "title":"New title from tests",
        "content": "New content from tests"
    }

    url = "http://{}:{}{}".format(host,port, route)

    rsp = request.post(url, data=json.dumps(data))

    if (rsp.status_code != 200):
        return -1
    return 0



def main():
    test01 = create_post_test()
    if(test01 == 0):
        print("test01 has passed")
    else:
        print("test01 has failed")

   


main()
