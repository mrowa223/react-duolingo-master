# React + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

# Backend

1. Configure ```.env``` file in the ```./backend``` folder
```conf
TG_BOT_TOKEN="YOUR_TELEGRAM_BOT_TOKEN"
GEMINI_API_KEY="YOUR_GEMINI_API_KEY"
```

2. Run commands:
```bash
cd ./backend
go run ./cmd/ -env-path .env
````