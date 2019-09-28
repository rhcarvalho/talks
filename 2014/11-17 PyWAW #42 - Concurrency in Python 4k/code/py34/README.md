# Python 3.4 concurrency example

This uses the `aiohttp` third-party library.

Create a virtual environment and install the dependencies with:

```
mktmpenv -p $(which python3.4) -r requirements.txt
cd -
```

Note: `mktmpenv` comes from `virtualenvwrapper`. There are many other ways to
install the dependencies, out of scope here.

Run it with:

```
python asyncio-get-request.py
```
