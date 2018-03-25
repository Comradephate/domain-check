[![Build status](https://badge.buildkite.com/d57a3a0fbe1914cd2833a0716d9af54d08daac8eb6e058f3f9.svg)](https://buildkite.com/jaron/domaincheck)

# domain-check

The idea is to make a website that allows users to log in and keep track of domains that they either own or are interested in, and be alerted when one has expired. It definitely won't work as quickly or as effectively as the literal dozens of registrar-run competitors for scooping up domains. :V

I'm just making this to learn Go.

# API

There's one route, `/api/v1/whois` which you can post json to in the form of `'{"name":"<somedomain.com>"}'` and it will respond with some information about that domain, also in json.