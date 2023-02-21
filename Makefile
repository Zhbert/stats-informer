run:
	werf compose up --docker-compose-command-options="-d"

stop:
	werf compose down

debug:
	werf compose up --follow