# Shion

[![Build Status](https://travis-ci.com/julioc98/shion.svg?token=4SjCRRz2dpNCgC3iccDx&branch=master)](https://travis-ci.com/julioc98/shion)

[Acesse a API](https://shion-api.herokuapp.com/)

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/julioc98/shion)


### Pré-requisitos

* [Golang](https://github.com/golang/go)

OU (Recomendado)

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)*

### Como rodar localmente?**

Baixe o repositório, entre no diretório e rode o comando:

```
make run/docker
```
Depois acesse a url
```
http://localhost:5001/
```

### Como rodar os testes?**

```
make test/docker
```

### Como fazer o deploy?

- Temos um Dockerfile para quando queremos levar essa aplicação para produção(Dockerfile.production). Ela é diferente da desnvolvimento porque ela usa `multi-stage build` para deixar a imagem bem menor com um S.O. mais leve e apenas o binário da aplicação para rodar.
- No caso estou fazendo o deploy no [Heroku](https://www.heroku.com/). Então só "commitar na master" ou apertar o botão de deploy:

  - [![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/julioc98/shion)


##### *Para Facilitar | **Com Docker + Docker Compose



