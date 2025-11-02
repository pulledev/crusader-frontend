# Crusader Backend
### Abstract
Das Crusader Backend ist der Versuch, verschiedene Prozesse zu etablieren, bestehende zu automatisieren und durch ein Datenbanksystem eine Redundanz-arme und Inkonsistenz-freie Speicherung umzusetzen. 

## Tech-Stack aka the nerdy Part

### Overview:
 - Frontend (Svelte 5)
 - Backend (GO & OpenAPI)
 - Datenbank (Postgres)
 - Backend auf Arma Server (C#) wird in Zukunft benutzt
 - Workflow and Automation (n8n)
 - Discord Bots (Kite)

### Frontend
 - Svelte 5, Svelte-Kit + TS
 - TailwindCSS, Bits-UI, Phosphor Icons
 - OpenAPI, SSO Discord
 - für CRUD-Operations mit cooler Datenvisualisierung
 - LLM Usage ist erlaubt

### Backend
 - GOLANG 1.24 (mby 1.25)
 - ORM: GORM
 - oapi-codegen
 - LLM Usage ist nicht gewollt

### Datenbank
 - Postgres

### Arma Server
 - vermutlich in C#
 - wird als externes System angesehen
 - vielleicht als Übergang GO Skript auf Server, welches Server Logs liest um Daten zu parsen und zu senden

### n8n
 - no Code System für literally alles
 - Anwendungsfall: Discord, Telegram, externe APIs für Automatisierung
   - Bei DC Änderung einfach anzupasen, ggf. zB von Mia, da no-Code
 - Nicht gedacht für: DB Zugriffe, Frontend->Backend etc. (siehe Communication)

### kite
 - no Code Discord Bots
 - Für DC User Interaktion
 - wenig Logik implementierbar
 - kann API Requests abschicken aber nicht durch Webhook getriggert werden

### Deployment
 - Docker Compose (Frontend, Backend, DB, n8n, kite)
 - FE und BE werden Dockerized

### Communication
<img src="techstack_diagram.png" style="width: 450px" alt="Techstack Diagramm">


# Development

Wir benutzen für VS-Code einen Docker Dev Container.
Dieser funktioniert nur mit VS-Code, aber nicht mit zB Goland(IntelliJ)

Folgendes ist vorausgesetzt für eine funktionieren Enwicklungsumgebung für Front- und Backend:
- Linux(zB über WSL)
- Go 1.25
- git, curl, postgres-client
- nodejs und npm
- Docker

Folgende Schritte sind zu beachten:

1. Datenbank und Adminer starten (im Root-Ordner)
    
    Abrufbar unter localhost:8081 und localhost:5432
    ```
   docker compose up -d
   ```
2. .env Datei mit Daten füllen
    ```
   DB_URL="host=localhost user=postgres password=crusader dbname=crusader port=5432 sslmode=disable TimeZone=Europe/Berlin"
   ```
3. In `crusader-back/` wechseln
4. Folgende Pakete installieren.
    ```
    go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/cweill/gotests/gotests@v1.6.0 \
    && go install golang.org/x/tools/gopls@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest
    ```
5. ORM Models migrieren
     ```
   go run migrate/migrate.go
   ```
6. API Backend kann gestartet werden (port 8080)
      ```
   go run .
   ```
7. In `crusader-front/` wechseln
8. Npm Pakete installieren
      ```
   npm i
   ```
9. Frontend Server starten (port 5173)
      ```
   npm run dev
   ```