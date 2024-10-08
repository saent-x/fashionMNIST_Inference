.PHONY: build start watch

build:
	NODE_ENV=production ./node_modules/.bin/postcss build tailwind.css -o public/styles/app.css

start:
	go run main.go

watch:
	NODE_ENV=development ./node_modules/.bin/postcss build ./public/styles/app.css -o public/styles/app.min.css -w
