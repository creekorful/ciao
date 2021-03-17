FROM scratch

ADD direktion /usr/bin/direktion

ENTRYPOINT ["/usr/bin/direktion"]