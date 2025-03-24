# CoursUT

[![test & build](https://github.com/tot0p/CoursUT/actions/workflows/test&build.yml/badge.svg?branch=main)](https://github.com/tot0p/CoursUT/actions/workflows/test&build.yml)
[![codecov](https://codecov.io/gh/tot0p/CoursUT/graph/badge.svg?token=DYEN01ALDK)](https://codecov.io/gh/tot0p/CoursUT)

## Sujet

### 1 Gestion d’un Parking Intelligent

Description : Un système permettant aux utilisateurs de réserver une place de parking dans un parking intelligent.
Fonctionnalités
1. Gestion des véhicules
   Ajouter un véhicule (plaque, type, durée de stationnement).
   Modifier/supprimer un véhicule enregistré.
   Vérifier si un véhicule est actuellement stationné.
2. Gestion des places
   Vérifier la disponibilité des places.
   Attribuer dynamiquement une place en fonction du type de véhicule.
   Gérer les réservations anticipées.
3. Calcul et gestion des tarifs
   Tarification en fonction du temps stationné et du type de véhicule.
   Tarifs dégressifs (ex: 1h = 5€, 2h = 9€, 3h = 12€).
   Paiement simulé avec validation du montant.
4. Génération et gestion des tickets
   Génération d’un ticket numérique avec QR Code.
   Affichage du temps restant


## Fonctionnalités implémentées

- [x] Gestion des véhicules
  - [x] Ajouter un véhicule (plaque, type, durée de stationnement).
  - [x] Modifier/supprimer un véhicule enregistré.
  - [ ] Vérifier si un véhicule est actuellement stationné.
- [x] Gestion des places
  - [ ] Vérifier la disponibilité des places.
  - [x] Attribuer dynamiquement une place en fonction du type de véhicule.
  - [x] Gérer les réservations anticipées.
- [x] Calcul et gestion des tarifs
  - [x] Tarification en fonction du temps stationné et du type de véhicule.
  - [x] Tarifs dégressifs (ex: 1h = 5€, 2h = 9€, 3h = 12€).
  - [ ] Paiement simulé avec validation du montant.
- [x] Génération et gestion des tickets
  - [x] Génération d’un ticket numérique avec QR Code.
  - [x] Affichage du temps restant

Ce projet est qu'une **API REST** qui permet de gérer un parking intelligent.
Développé en [golang](https://golang.org/).
Il utilise le framework [fiber](https://gofiber.io/) pour la gestion des routes.

## Run it

### Local

#### Prérequis

- [golang](https://golang.org/)
- [gcc](https://gcc.gnu.org/)

#### Run tests

```bash
go test ./...
```

or verbose mode

```bash
go test -v ./...
```

#### Compilation

```bash
go build -o main cmd/App/main.go
```

#### Lancement

```bash
./main
```

### Docker

#### Build image

```bash
docker build -t coursut .
```

#### Run container

```bash
docker run -p 8080:8080 coursut
```

## Authors

- [tot0p](https://github.com/tot0p)
- [mkarten](https://github.com/mkarten)