```
mon-projet/
│
├── cmd/                # Entrées principales de l’application (exécutables)
│   └── mon-app/
│       └── main.go
│
├── pkg/                # Packages publics réutilisables par d'autres projets
│   └── monmodule/
│       └── ...
│
├── internal/           # Code privé à ce projet (non importable de l’extérieur)
│   └── monmodule/
│       └── ...
│
├── api/                # Définitions d’API (OpenAPI, Protobuf, etc.)
│
├── configs/            # Fichiers de configuration (YAML, JSON…)
│
├── scripts/            # Scripts d’automatisation (build, déploiement…)
│
├── test/               # Données ou helpers spécifiques aux tests
│
├── go.mod              # Fichier de module Go
├── go.sum              # Sommes de contrôle des dépendances
└── README.md           # Documentation du projet
```