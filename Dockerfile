FROM scratch

ADD ciao /usr/bin/ciao

ENTRYPOINT ["/usr/bin/ciao"]