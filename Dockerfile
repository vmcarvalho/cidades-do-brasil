FROM golang:latest

# parameters
ARG DATABASE_URL
ENV DATABASE_URL ${DATABASE_URL}

ARG CSV_DATABASE_URL=ftp://geoftp.ibge.gov.br/cartas_e_mapas/bases_cartograficas_continuas/bc250/versao2017/lista_de_nomes_geograficos/bc250_nomesgeograficos.csv;

# for internal use
ARG TEMP_DOWNLOAD_PATH=/tmp/input.csv
ENV TEMP_DOWNLOAD_PATH ${TEMP_DOWNLOAD_PATH}
ARG APP_DIR=$GOPATH/src/github.com/vmcarvalho/cidades-do-brasil/

# download input file
RUN wget --output-document $TEMP_DOWNLOAD_PATH $CSV_DATABASE_URL

ADD ./ $APP_DIR

# Get dependencies
RUN go get github.com/globalsign/mgo;

WORKDIR $APP_DIR
RUN go build -o main
CMD ./main ${TEMP_DOWNLOAD_PATH} ${DATABASE_URL}