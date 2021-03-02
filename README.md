## Sobre

Projeto de avaliacao sobre a linguagem GO.


## Softwares necessários

 * GO <= 1.15.8
 * PostGres 


## Instalando e configurando

Clone este repositório para um servidor web.

https://github.com/wrequiao/avaliacao.git

Configure no arquivo config/dev_config.json os dados para conectar no banco de dados.
rodar os comandos dentro do diretório do projeto
- go get github.com/lib/pq
- go get github.com/tkanos/gonfig

Execute os scripts abaixo no Postgres para criar o banco de dados e as suas tabelas.

CREATE DATABASE db_arquivo;

CREATE TABLE arquivo
(
   id                SERIAL PRIMARY KEY,
   nome              VARCHAR(200) NOT NULL,
   datacriacao       DATE NOT NULL,
   dataprocessamento DATE NULL
); 


CREATE TABLE arquivolinha
(
   id        SERIAL PRIMARY KEY,
   linha     VARCHAR(2000) NOT NULL,
   idarquivo INT NOT NULL,
   FOREIGN KEY (idarquivo) REFERENCES arquivo (id)
); 

CREATE TABLE dados
  (
     id                 SERIAL PRIMARY KEY,
     cpf                VARCHAR(20) NULL,
     private            VARCHAR(1) NULL,
     incompleto         VARCHAR(1) NULL,
     dataultimacompra   VARCHAR(20) NULL,
     ticketmedio        VARCHAR(20) NULL,
     ticketultimacompra VARCHAR(20) NULL,
     lojamaisfrequente  VARCHAR(20) NULL,
     lojadaultimacompra VARCHAR(20) NULL,
     status             VARCHAR(2000) NULL,
     idarquivo          INT NOT NULL,
     FOREIGN KEY (idarquivo) REFERENCES arquivo (id)
  ); 

O projeto esta configurado para rodar na porta 8080
http://localhost:8080

## Fluxo de funcionamento

- O arquivo deve ser feito upload.
- Logo em seguida o arquivo deve ser processado.
- Após o processamento é possível visualizar os dados importados e o seu status.
