# Greenlight API

**REST API learning project built with Go and PostgreSQL.**

Greenlight est un projet backend réalisé pour approfondir la conception d’API HTTP robustes avec **Go**, en mettant l’accent sur la structure du code, PostgreSQL, les migrations, la validation et la gestion propre des erreurs.

> 📚 Projet d’apprentissage orienté Backend Engineering.

---

## 🎯 Objectif

L’objectif principal est de construire une API REST proche d’un environnement de production en utilisant autant que possible les outils standards de l’écosystème Go.

Le projet permet notamment de travailler sur :

* la conception d’API HTTP ;
* la structuration d’une application Go ;
* PostgreSQL ;
* les migrations SQL ;
* la validation ;
* les middlewares ;
* la gestion des erreurs ;
* la configuration ;
* la gestion propre du cycle de vie du serveur.

---

## 🛠 Stack

* Go
* PostgreSQL
* SQL
* `net/http`
* Database migrations
* Makefile

---

## 🏗 Structure

```text
.
├── cmd/
│   └── api/
│
├── internal/
│
├── migrations/
│
├── Makefile
├── go.mod
└── go.sum
```

### `cmd/api`

Contient le point d’entrée de l’API et la configuration principale de l’application.

### `internal`

Contient les composants internes de l’application, notamment la logique partagée et l’accès aux données.

### `migrations`

Contient les migrations SQL permettant de versionner le schéma PostgreSQL.

---

## 🧠 Concepts travaillés

### HTTP & REST

* routing ;
* handlers HTTP ;
* encodage et décodage JSON ;
* codes de statut HTTP ;
* erreurs API cohérentes.

### Middleware

* logging ;
* récupération après panic ;
* gestion des en-têtes ;
* contrôle du cycle des requêtes.

### Validation

Les entrées reçues par l’API sont validées avant leur traitement afin de conserver un modèle de données cohérent.

### PostgreSQL

Le projet utilise PostgreSQL comme base relationnelle.

Les problématiques étudiées comprennent notamment :

* requêtes SQL ;
* connexion à la base ;
* gestion des erreurs SQL ;
* timeouts ;
* migrations.

---

## 🗄 Migrations

Le schéma de base de données est versionné à l’aide de migrations SQL.

```text
migrations/
├── 000001_...
├── 000002_...
└── ...
```

Cette approche permet de reproduire et faire évoluer le schéma de façon contrôlée.

---

## ⚙️ Configuration

L’application utilise une configuration externe pour les éléments dépendant de l’environnement.

Exemples :

```text
Database DSN
Server port
Environment
Connection pool settings
```

Les secrets ne doivent jamais être intégrés directement au code source.

---

## 🚀 Installation

### 1. Cloner le repository

```bash
git clone https://github.com/johnmabs/greenlight.git
cd greenlight
```

### 2. Installer les dépendances

```bash
go mod download
```

### 3. Configurer PostgreSQL

Créer une base de données puis fournir la chaîne de connexion utilisée par l’application.

### 4. Exécuter les migrations

Appliquer les migrations présentes dans :

```text
migrations/
```

### 5. Lancer l’API

```bash
go run ./cmd/api
```

---

## 🔨 Makefile

Le repository contient un `Makefile` afin de centraliser les commandes courantes du projet.

Selon l’état actuel du projet, il peut notamment être utilisé pour :

```bash
make run
make audit
make build
```

Consulter le `Makefile` pour connaître les commandes disponibles.

---

## 🧪 Qualité

Le projet cherche à conserver une base backend propre grâce à :

* formatage Go ;
* analyse statique ;
* dépendances explicites ;
* séparation des responsabilités ;
* gestion centralisée des erreurs ;
* configuration externe ;
* migrations versionnées.

---

## 📚 Pourquoi Go ?

Ce projet me permet d’explorer une approche backend différente de ma stack principale PHP/Symfony.

Go apporte notamment une approche intéressante autour de :

* simplicité du langage ;
* compilation ;
* concurrence ;
* tooling intégré ;
* services HTTP performants ;
* faible coût d’exécution.

---

## 📈 Évolutions possibles

* authentification ;
* rate limiting ;
* tests d’intégration ;
* métriques ;
* structured logging ;
* graceful shutdown avancé ;
* pagination ;
* filtering ;
* background jobs ;
* CI avec GitHub Actions.

---

## 👨‍💻 Auteur

**John Mabiala**

Full-Stack / Backend Developer

* [GitHub](https://github.com/johnmabs)
* [LinkedIn](https://linkedin.com/in/john-mabiala)

---

## 📌 Portfolio

Ce repository fait partie de mon portfolio Backend Engineering.

Il illustre mon apprentissage de **Go, PostgreSQL et de la conception d’API HTTP**, en complément de mes projets principaux développés avec Symfony et Next.js.
