# Python 2.7 concurrency example

This uses the `gevent` third-party library.

Create a virtual environment and install the dependencies with:

```
mktmpenv -p $(which python2) -r requirements.txt
cd -
```

Note: `mktmpenv` comes from `virtualenvwrapper`. There are many other ways to
install the dependencies, out of scope here.

Run it with:

```
python gevent-get-request.py
```
