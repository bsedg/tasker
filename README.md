# tasker

[ ![Codeship Status for bsedg/irest](https://codeship.com/projects/1f8bd7f0-598d-0134-a57b-0ac50249f2fe/status?branch=master)](https://codeship.com/projects/173070)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsedg/tasker)](https://goreportcard.com/report/github.com/bsedg/tasker)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Tasker is a service to manage tasks that can be scheduled.

## Development

```
# build and run the service
docker-compose up -d --force-recreate taskservice

# initialize the database by creating tables, etc.
curl -i -x POST localhost:80/db/init --header "X-Tasker-Authentication: <AUTH_KEY>"

# create a task
curl -i -X POST localhost:80/tasks \
    -d '{"name": "test", "action": "noop", "time": "now"}'

# get all tasks
curl -i localhost:80/tasks
```

### Integration tests
Using the iREST framework to create HTTP integration tests, https://github.com/bsedg/irest, integration tests can be run with the appropriate build tags or with the docker-compose test.

`docker-compose -f docker-compose.test.yml -p ci up -d`
