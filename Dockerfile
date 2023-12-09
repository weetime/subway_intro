FROM 139.198.4.111:30002/library/alpine:3.18
RUN mkdir -p /apps/logs
COPY ./build/ /apps/
COPY ./config/ /apps/config/