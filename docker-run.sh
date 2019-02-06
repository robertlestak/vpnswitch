docker run --rm -it \
  --net=host \
  --privileged \
  --env-file .env \
  -v $PWD/data:/data \
  --name vpn vpn
