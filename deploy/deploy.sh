#! /bin/sh
pjName='maxmicro'
pjGitName='maxmicro'

if [ ! -d "www/${pjGitName}" ]; then
    echo 'git clone'
    # shellcheck disable=SC2164
    cd www
    git clone https://liuzhaomax@github.com/liuzhaomax/${pjGitName}.git
    # shellcheck disable=SC2164
    cd ${pjGitName}
else
    echo 'git pull'
    # shellcheck disable=SC2164
    cd www/${pjGitName}
    git pull
fi

docker stop ${pjName}
docker rm -f ${pjName}
docker rmi -f ${pjName}-i

docker build -t ${pjName}-i .
docker run --name=${pjName} -d --restart=always --privileged=true --network my_net -p 8000:8000 -v /home/liuzhaomax1/ftp/static:/usr/src/app/static ${pjName}-i