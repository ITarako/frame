include .env

build: git-pull check-env migrate-up build-image docker-up

git-pull:
	docker run -it --rm -v="${SSH_KEYS}:/root/.ssh" -v="${PROJECT_DIR}:/opt/project" --add-host=${GITLAB_HOST} ${DOCKER_ID}/git git clone ${PROJECT_REPO} /opt/project
	docker run -it --rm -v="${SSH_KEYS}:/root/.ssh" -v="${PROJECT_DIR}:/opt/project" --add-host=${GITLAB_HOST} -w /opt/project ${DOCKER_ID}/git git checkout ${PROJECT_BRANCH}
	docker run -it --rm -v="${SSH_KEYS}:/root/.ssh" -v="${PROJECT_DIR}:/opt/project" --add-host=${GITLAB_HOST} -w /opt/project ${DOCKER_ID}/git git pull origin ${PROJECT_BRANCH}

check-env:
ifeq ($(wildcard ${PROJECT_DIR}/.env),)
		cp ${PROJECT_DIR}/.env.example ${PROJECT_DIR}/.env
		@echo "Заполните ${PROJECT_DIR}/.env"
		exit 1
endif

build-image:
	docker build --tag=baseapp_api ${PROJECT_DIR}

build-image-no-cache:
	docker build --no-cache --tag=baseapp_api ${PROJECT_DIR}

migrate-up:
	docker-compose --env-file ${PROJECT_DIR}/.env -f ${PROJECT_DIR}/deploy/docker-compose.yml --profile build run --rm migrate up

docker-up:
	docker-compose --env-file ${PROJECT_DIR}/.env -f ${PROJECT_DIR}/deploy/docker-compose.yml up -d
	#docker-compose --env-file ${PROJECT_DIR}/.env -f ${PROJECT_DIR}/deploy/docker-compose.yml -f docker.yml up -d