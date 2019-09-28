#!/usr/bin/env python2.7
import gevent
from gevent import monkey
monkey.patch_all()
import requests


def main():
    urls = [
        "https://google.fr",
        "https://google.hu",
        "https://google.pl",
    ]
    jobs = [gevent.spawn(download, url) for url in urls]
    gevent.joinall(jobs, timeout=2)


def download(url):
    response = requests.get(url)
    status = response.status_code
    print("{}: {} {} ({} bytes)".format(url, status, response.reason, len(response.content)))


if __name__ == "__main__":
    main()
