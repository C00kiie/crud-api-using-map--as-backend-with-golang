#!/usr/bin/env python3
import requests as request

# test01
# route: /
# method: GET
# expected result: returns 200


# test02
# route: /post/{id}
# method: GET
# expected result: returns 200


host = "127.0.0.1"
port = "8080"



def get_all_posts():
    route = '/'
    url = "http://{}:{}{}".format(host,port, route)

    if(request.get(url).status_code != 200):
        return -1
    return 0

def get_post_by_id():
    route = '/post/f0668f75-b99e-447f-8495-fbe660bb1892'
    url = "http://{}:{}{}".format(host,port, route)
    if(request.get(url).status_code != 200):
        return -1
    return 0


def main():
    test01 = get_all_posts()
    if(test01 == 0):
        print("test01 has passed")
    else:
        print("test01 has failed")

    test02 = get_post_by_id()
    if(test02 == 0):
        print("test01 has passed")
    else:
        print("test01 has failed")


main()
