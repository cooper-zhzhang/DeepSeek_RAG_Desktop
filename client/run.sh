#/usr/bin/env bash

docker pull qdrant/qdrant
docker run -itd --name qdrant -p 6333:6333 qdrant/qdrant
