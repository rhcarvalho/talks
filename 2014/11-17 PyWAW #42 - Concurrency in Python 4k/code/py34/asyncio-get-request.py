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


# *** Notes ***
# - The `@asyncio.coroutine` decorator is optional but recommended
# - It is hard to tell whether things are really happening concurrently or
#   sequentially. From the call site, it is not obvious that calls to
#   `download(url)` run a coroutine...
# - Code is difficult to read, it doesn't look like "normal Python"
# - `yield from` is not a concept everybody coming to the language can quickly
#   understand
# - Same syntax `def ...`, strickingly different behavior depending on the
#   presence of a keywork in the body!
# - Very hard to debug when things go wrong
# - More complicated to chain calls when compared to regular functions
# - Coroutines infect the APIs, everything needs to be specially crafted to
#   cooperate
# - Inconsistent syntax, sometimes parentheses are necessary surrounding
#   `yield from` expressions, but not always
