FROM alpine:3.7
RUN adduser -S jean
ENV GOPATH /root/go
RUN apk update && \
    apk add go \
            git \
            tzdata \
            glide \
            make \
            musl-dev \
    && rm -rf /var/cache/apk/* \
              /tmp/* \
              $HOME/.cache

COPY . ${GOPATH}/src/github.com/xorilog/jean-bobot
WORKDIR ${GOPATH}/src/github.com/xorilog/jean-bobot
RUN mkdir -p ${GOPATH}/{src,bin} \
    && make install \
    && make build \
    && chmod +x build/jean-bobot \
    && mv build/jean-bobot /usr/local/bin/ \
    && rm -rf ${GOPATH}
USER jean
ENTRYPOINT ["jean-bobot"]
