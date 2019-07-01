FROM alpine:3.2

ADD conf /conf
ADD html /html
ADD homeweb-web /homeweb-web

WORKDIR /

ENTRYPOINT [ "/homeweb-web" ]

EXPOSE 8888
