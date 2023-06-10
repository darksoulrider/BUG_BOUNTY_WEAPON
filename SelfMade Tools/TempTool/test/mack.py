import requests


host = "https://www.facebook.com"

data = requests.get(host)


with open("index.html","w", encoding='utf-8') as f:
    f.write(data.text)
