#docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
cwd=$(PWD)
all:
	java -jar antlr-4.13.1-complete.jar -Dlanguage=Go -o parser grammar/Tafexpr.g4
evaparser:
	java -jar antlr-4.13.1-complete.jar -Dlanguage=Go -o parser grammar/Evalang.g4