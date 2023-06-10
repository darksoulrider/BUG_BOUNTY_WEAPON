import requests

url = "https://maps2.zomato.com/osm/10/732/427.png"

data = requests.get(url)

print(data.status_code


)


# https://blend.com/resources/?format=   -> check xss and  all
