import aiohttp
import asyncio

@asyncio.coroutine
def download(uri):
    response = yield from aiohttp.request('GET', uri)
    return (yield from response.read_and_close())

def download2(uri):
    response = yield from aiohttp.request('GET', uri)
    return (yield from response.read_and_close())

@asyncio.coroutine
def print_beginning(uri):
    result = yield from download(uri)
    print(result[:100])

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait([
        print_beginning('http://google.pl'),
        print_beginning('http://google.fr'),
        print_beginning('http://google.hu'),
    ]))

print(download, download2)