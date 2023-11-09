If you want to remove all the dangling docker images: https://nickjanetakis.com/blog/docker-tip-31-how-to-remove-dangling-docker-images
docker rmi -f $(docker images -f "dangling=true" -q)

