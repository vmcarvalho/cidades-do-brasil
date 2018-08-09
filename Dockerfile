FROM golang:latest
RUN mkdir /internal_data; cd /internal_data; \
wget ftp://geoftp.ibge.gov.br/cartas_e_mapas/bases_cartograficas_continuas/bc250/versao2017/lista_de_nomes_geograficos/bc250_nomesgeograficos.csv;
CMD ["sleep", "infinity"]