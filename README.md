# Kahvi- ja tee API

Kesätyöhaun harjoitustyö. API, joka tarjoaa mahdollisuuden hakea kahvi- ja teetietoja sekä lisätä uusia tietoja.
 
## Toteutus

Tietokanta on toteutettu MongoDB:llä (MongoDB Atlas), backend Go:lla ja frontend React + TS combolla.

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"> <img src="https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black"> <img src="https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white"> <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white">

# Käyttö, käynnistys & testaus

Alla olevat ohjeet on testattu Linuxilla (Fedora 37) ja MacOS:lla (Ventura 13.1). Windowsilla ajoa ei ole testattu, mutta toiminnalisuudessa ei pitäisi olla ongelmia (Esim. Git BASH:lla ajaessa). Jos kuitenkin ongelmia ilmenee, ottakaa yhteyttä, jotta saadaan homma toimimaan.

## 0. Tarvittavat työkalut 

Jotta projekti toimisi, tarvitset seuraavat ohjelmat/työkalut asennettuna:

* Go (backend): https://golang.org/doc/install
* Node.js: https://nodejs.org/en/download/
* NPM: https://www.npmjs.com/get-npm

## 1. Lataa projekti

1. Lataa projekti komennolla `git clone https://github.com/emilsto/alma-task` 
2. Siirry projektin juurihakemistoon (`cd alma-task`)

## 2. env tiedostot

.env tiedostot ovat gitignoressa, joten ne pitää luoda itse. Helpoin tapa on ajaa init_env.sh scripti, joka luo .env tiedostot ja asettaa tarvittavat arvot (Löytyy /backend -hakemistosta). HUOM. Scriptin ajo vaati tietysti bash shellin. Lähettämässäni sähköpostissa MONGO_URI:n arvo, joka tulee lisätä .env tiedostoon, vaikka scriptiä käyttäisikin.

### 2.1. Jos haluat tehdä .env tiedostot itse, niin seuraavat ohjeet ovat sinulle.

1. Luo .env tiedosto backendin juurihakemistoon (/backend)
2. Lisää .env tiedostoon seuraavat rivit:
```env
PORT = ":5001"
MONGO_URI = "ks. sähköposti"
MONGO_COLLECTION = "beverages"
```
3. Vaihda MONGO_URI:n arvo sähköpostissa saamasi arvoon

4. Luo .env tiedosto testihakemistoon (/backend/tests)
5. Lisää .env tiedostoon seuraavat rivit:
```env
PORT = ":5001"
MONGO_URI = "ks. sähköposti"
MONGO_COLLECTION = "test_beverages"
```
6. Vaihda MONGO_URI:n arvo sähköpostissa saamasi arvoon


## 3. Backend

1. Siirry backendin juurihakemistoon (`/backend`)
2. Asenna riippuvuudet komennolla `go mod download`
3. Käynnistä backend komennolla `go run main.go` TAI `go build main.go && ./main`, jos haluat käynnistää backendin buildattuna.

Selaimen pitäisi näyttää osoitteessa `http://localhost:5001/api/v1/beverages` jotain tämän kaltaista:

```json
[
  {
    "id": "63dac1047505861c90b4c454",
    "name": "Ethiopian Harrar",
    "weight": 500,
    "price": 12.99,
    "roast_degree": 3,
    "type": "Coffee"
  },
  {
    "id": "63dacb39d0feac7ab673c465",
    "name": "Juhlamokka",
    "weight": 500,
    "price": 6.5,
    "roast_degree": 2,
    "type": "Coffee"
  },
  ...
]
```

## 4. Frontend

1. Siirry frontendin juurihakemistoon (`/frontend`)
2. Asenna riippuvuudet komennolla `npm install`
3. Käynnistä frontend komennolla `npm start`

Selaimen pitäisi avautua osoitteessa `http://localhost:3000/` ja näyttää kahvi- ja teetietoja.

## 5. Testaus

Backendin testit on toteutettu Go:n testauskirjastolla. Testit ajetaan siirtymällä backend/tests -hakemistoon ja ajamalla komento `go test`.
Onnistuneiden testien tulisi näyttää tältä:

```bash
$ go test

PASS
Connected to MongoDB
Collection name:  testing_beverages
{ObjectID("000000000000000000000000") Pirkka Costa Rica 500 5.5 0 Coffee}
PASS
ok      github.com/emilsto/alma-task/tests      0.640s
```

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

