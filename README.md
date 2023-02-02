# Kahvi- ja teetietokanta

Kesätyöhaun harjoitustyö. Tietokanta, johon voi lisätä kahvi- ja teetietoja. Tietokantaan voi lisätä kahvin tai teen nimen, maan, laadun, hinnan, maun ja muiden ominaisuuksien. Tietokannasta voi hakea kaikki kahvi- ja teetiedot tai hakusanalla kahvi- ja teetiedot, joiden nimi sisältää hakusanan.

## Toteutus

Tietokanta on toteutettu MongoDB:llä (MongoDB Atlas), backend Go:lla ja frontend React + TS comboilla.

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"> <img src="https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black"> <img src="https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white"> <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white">

# How to run

Alla olevat ohjeet on testattu Linuxilla (Fedora 37) ja MacOS:lla (Big Sur 11.4). Windowsilla ajoa ei ole testattu, mutta toiminnalisuudessa ei pitäisi olla ongelmia. Jos kuitenkin ongelmia ilmenee, ottakaa yhteyttä, niin korjaan ohjeet!

## Prerequisites

* Go (backend): https://golang.org/doc/install
* Node.js: https://nodejs.org/en/download/
* NPM: https://www.npmjs.com/get-npm


## Backend

1. Luo .env tiedosto backendin juurihakemistoon (/backend)
2. Lisää .env tiedostoon seuraavat rivit:
```env
PORT = 5000
MONGO_URI = "ks. sähköposti"
MONGO_COLLECTION = "beverages"
```

3. Käynnistä backend komennolla `go run main.go` TAI `go build main.go` ja `./main`)

## Frontend

1. Käynnistä frontend komennolla `npm start`

# Tilanne & TODO

### Backend

- [x] Hae kaikki kahvi- ja teetiedot
- [x] Hae hakusanalla kahvi- ja teetiedot, joiden nimi sisältää hakusanan (Tyyliin ILIKE %haettava% sql:ssä).
- [x] Luo uusi kahvi- tai teetieto
- [ ] Testit
    - [ ] Testaa reitit
    - [ ] Testaa tietokantayhteydet
    - [ ] Testaa tietokantatoiminnot

### Frontend

- [x] Hae kaikki kahvi- ja teetiedot
- [ ] Hae hakusanalla kahvi- ja teetiedot, joiden nimi sisältää hakusanan (Tyyliin ILIKE %haettava% sql:ssä).
- [ ] Luo uusi kahvi- tai teetieto

# Dependencies

## Backend

* Router: https://pkg.go.dev/github.com/gorilla/mux
* Env: https://pkg.go.dev/github.com/joho/godotenv
* MongoDB: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
* CORS: https://pkg.go.dev/github.com/rs/cors

## Frontend

* React: https://reactjs.org/
* Typescript: https://www.typescriptlang.org/

