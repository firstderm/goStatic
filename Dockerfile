FROM scratch

COPY goStatic /
ENTRYPOINT ["/goStatic"]