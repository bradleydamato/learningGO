import sys, json;
data = json.load(sys.stdin)
for movie in data:
	print(movie)


