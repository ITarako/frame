#!/bin/bash

DOCKER_ID='freematiq'
PROJECT_REPO='git@gitlab.freematiq.com:quartz/yii2-base-app.git'
PROJECT_BRANCH='master'
PROJECT_DIR='/opt/baseapp/project'
SSH_KEYS='/home/users/.ssh'
GITLAB_HOST='gitlab.freematiq.com:46.4.155.93'

SERVER_NAME='baseapp.local'

docker run -it --rm -v="$SSH_KEYS:/root/.ssh" -v="$PROJECT_DIR:/opt/project" --add-host=$GITLAB_HOST $DOCKER_ID/git git clone $PROJECT_REPO /opt/project
docker run -it --rm -v="$SSH_KEYS:/root/.ssh" -v="$PROJECT_DIR:/opt/project" --add-host=$GITLAB_HOST -w /opt/project $DOCKER_ID/git git checkout $PROJECT_BRANCH
docker run -it --rm -v="$SSH_KEYS:/root/.ssh" -v="$PROJECT_DIR:/opt/project" --add-host=$GITLAB_HOST -w /opt/project $DOCKER_ID/git git pull origin $PROJECT_BRANCH

if [ ! -e $PROJECT_DIR/.env ]; then
    cp $PROJECT_DIR/.env.example $PROJECT_DIR/.env;
    echo "Заполните $PROJECT_DIR/.env";
    exit 0
fi

source $PROJECT_DIR/.env

docker-compose -f $PROJECT_DIR/deploy/docker-compose.yml --profile build run --rm migrate up
docker-compose -f $PROJECT_DIR/deploy/docker-compose.yml up -d
#docker-compose -f $PROJECT_DIR/deploy/docker-compose.yml -f docker.yml up -d
