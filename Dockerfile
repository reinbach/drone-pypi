FROM alpine
ADD pypi /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/pypi
