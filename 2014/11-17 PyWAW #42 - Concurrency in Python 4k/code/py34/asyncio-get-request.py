#!/usr/bin/env python3.4
import asyncio

import aiohttp


def main():
    urls = [
        "https://google.fr",
        "https://google.hu",
        "https://google.pl",
    ]
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait([
        print_info(url) for url in urls
    ]))


@asyncio.coroutine
def download(url):
    response = yield from aiohttp.request("GET", url)
    return (yield from response.read_and_close())


@asyncio.coroutine
def print_info(url):
    result = yield from download(url)
    print("{}: {} ({} bytes)".format(url, "???", len(result)))


if __name__ == "__main__":
    main()
