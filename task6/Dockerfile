FROM golang:1.9-stretch
LABEL maintainer="Muh. Tarmizi <muh.tarmizi@koki.id>" 
ENV KOKI_NODE_NO 1 
RUN mkdir -p "/usr/bin/koki" 
WORKDIR /usr/bin/koki 
CMD ["./api", "-c", "config.yml"]