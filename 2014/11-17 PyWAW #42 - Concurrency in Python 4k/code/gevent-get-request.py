import gevent
from gevent import monkey
mokey.patch_all()
import requests

def download(uri):
    response = requests.get(uri)
    return response.content

def print_beginning(uri):
    result = download(uri)
    print(result[:100])

if __name__ == '__main__':
    urls = ['http://google.pl', 'http://google.fr', 'http://google.hu']
    jobs = [gevent.spawn(socket.gethostbyname, url) for url in urls]
    gevent.joinall(jobs, timeout=2)
