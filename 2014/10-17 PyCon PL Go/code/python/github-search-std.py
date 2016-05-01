#!/usr/bin/env python
from __future__ import print_function
# // begin OMIT
import urllib2, json

query, language = "python", "go"

url = ("https://api.github.com/search/repositories?q={}"
       "+language:{}&sort=stars&order=desc&per_page=3").format(query, language)
resp = urllib2.urlopen(url)
data = json.load(resp)

print(" #  {:<20} {}".format("Repo URL", "Stars"))
print("-" * 30)
for i, r in enumerate(data["items"], 1):
    print("{:02d}. {:<20} {:d}".format(i, r["full_name"], r["stargazers_count"]))
