NAME=golang-react-template

all: ui/build ${NAME}

${NAME}: ui/build
	go build -o ${NAME}

ui/build:
	cd ui; \
	npm install; \
	npm run build
