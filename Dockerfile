FROM python:2-alpine
ADD drone-pypi /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone-pypi
