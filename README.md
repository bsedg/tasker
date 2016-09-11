# tasker

[ ![Codeship Status for bsedg/irest](https://codeship.com/projects/1f8bd7f0-598d-0134-a57b-0ac50249f2fe/status?branch=master)](https://codeship.com/projects/173070)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsedg/tasker)](https://goreportcard.com/report/github.com/bsedg/tasker)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Tasker is a service to manage tasks that can be scheduled.

## Development

```
# build the service
docker-compose build

# run the service
docker-compose up

# assuming 'dockerhost' for the host
# create a task
curl -i -X POST dockerhost:8080/tasks \
    -d '{"name": "test", "action": "noop", "time": "now"}'

# get all tasks
curl -i dockerhost:8080/tasks
```
