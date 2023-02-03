# Kahvi- ja tee API

Kesätyöhaun harjoitustyö. API, joka tarjoaa mahdollisuuden hakea kahvi- ja teetietoja sekä lisätä uusia tietoja.
 
## Toteutus

Tietokanta on toteutettu MongoDB:llä (MongoDB Atlas), backend Go:lla ja frontend React + TS combolla.

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"> <img src="https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black"> <img src="https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white"> <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white">

# Käyttö, käynnistys & testaus

Alla olevat ohjeet on testattu Linuxilla (Fedora 37) ja MacOS:lla (Ventura 13.1). Windowsilla ajoa ei ole testattu, mutta toiminnalisuudessa ei pitäisi olla ongelmia (Esim. Git BASH:lla ajaessa). Jos kuitenkin ongelmia ilmenee, ottakaa yhteyttä, niin korjaan ohjeet!

## Prerequisites

* Go (backend): https://golang.org/doc/install
* Node.js: https://nodejs.org/en/download/
* NPM: https://www.npmjs.com/get-npm

## env tiedostot

.env tiedostot ovat gitignoressa, joten ne pitää luoda itse. Helpoin tapa on ajaa init_env.sh scripti, joka luo .env tiedostot ja asettaa tarvittavat arvot (Löytyy /backend -hakemistosta). Jos haluat tehdä .env tiedostot itse, niin seuraavat ohjeet ovat sinulle.

1. Luo .env tiedosto backendin juurihakemistoon (/backend)
2. Lisää .env tiedostoon seuraavat rivit:
```env
PORT = 5000
MONGO_URI = "ks. sähköposti"
MONGO_COLLECTION = "beverages"
```

3. Luo .env tiedosto testien juurihakemistoon (/backend/tests)
4. Lisää .env tiedostoon seuraavat rivit:
```env
MONGO_URI = "ks. sähköposti"
MONGO_COLLECTION = "test_beverages"
```


## Backend

1. Siirry backendin juurihakemistoon (`cd backend`)
2. Asenna riippuvuudet komennolla `go mod download`
3. Käynnistä backend komennolla `go run main.go` TAI `go build main.go && ./main`, jos haluat käynnistää backendin buildattuna.

## Frontend

1. Siirry frontendin juurihakemistoon (`cd frontend`)
2. Asenna riippuvuudet komennolla `npm install`
3. Käynnistä frontend komennolla `npm start`

# Tilanne & TODO

### Backend

- [x] Hae kaikki kahvi- ja teetiedot
- [x] Hae hakusanalla kahvi- ja teetiedot, joiden nimi sisältää hakusanan (Tyyliin ILIKE %haettava% sql:ssä).
- [x] Luo uusi kahvi- tai teetieto
- [x] Testit
    - [x] Testaa reitit
    - [x] Testaa tietokantayhteydet
    - [x] Testaa tietokantatoiminnot

### Frontend

- [x] Hae kaikki kahvi- ja teetiedot
- [x] Hae hakusanalla kahvi- ja teetiedot, joiden nimi sisältää hakusanan (Tyyliin ILIKE %haettava% sql:ssä).
- [x] Luo uusi kahvi- tai teetieto

# Dependencies

## Backend

* Router: https://pkg.go.dev/github.com/gorilla/mux
* Env: https://pkg.go.dev/github.com/joho/godotenv
* MongoDB: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
* CORS: https://pkg.go.dev/github.com/rs/cors

## Frontend

* React: https://reactjs.org/
* Typescript: https://www.typescriptlang.org/

