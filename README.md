# CoursUT

[![test & build](https://github.com/tot0p/CoursUT/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/tot0p/CoursUT/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/tot0p/CoursUT/graph/badge.svg?token=DYEN01ALDK)](https://codecov.io/gh/tot0p/CoursUT)

## 1 Gestion d’un Parking Intelligent

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