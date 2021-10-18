FROM goacme/lego
WORKDIR /certs
VOLUME [ "/certs" ]
ENV DNSPOD_API_KEY ""
ADD main /main
ENTRYPOINT [ "/main" ]