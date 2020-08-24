FROM nginx

MAINTAINER lanys "550709871@qq.com"

WORKDIR /src/godocker

ADD api /src/godocker

EXPOSE 9090
ENTRYPOINT ["./user", "-f", "./etc/config.json"]