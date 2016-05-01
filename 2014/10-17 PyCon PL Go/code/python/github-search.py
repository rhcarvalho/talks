#!/usr/bin/env python
from __future__ import print_function
# // begin OMIT
import requests // HL

query, language = "python", "go"

url = ("https://api.github.com/search/repositories?q={}"
       "+language:{}&sort=stars&order=desc&per_page=3").format(query, language)
resp = requests.get(url) // HL
data = resp.json() // HL

print(" #  {:<20} {}".format("Repo URL", "Stars"))
print("-" * 30)
for i, r in enumerate(data["items"], 1):
    print("{:02d}. {:<20} {:d}".format(i, r["full_name"], r["stargazers_count"]))
