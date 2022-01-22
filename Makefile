build:
	docker build -t cal-tracker:latest .

up-mongo:
	docker-compose -f mongo-compose.yaml up

up:
	docker-compose up