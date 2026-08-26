<div align="center">

# 🤖 Discord Bot en Go

**Un bot multipropósito para Discord rápido, eficiente y escalable, desarrollado con Go y DiscordGo.**

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Discord](https://img.shields.io/badge/DiscordGo-v0.28+-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://github.com/bwmarrin/discordgo)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-En_Desarrollo-orange?style=for-the-badge)]()

<p align="center">
  <a href="#-características-actuales">Características</a> •
  <a href="#-roadmap--futuras-implementaciones">Roadmap</a> •
  <a href="#-instalación-y-despliegue">Instalación</a> •
  <a href="#-configuración">Configuración</a> •
  <a href="#-tecnologías">Tecnologías</a>
</p>

---

</div>

## 📌 Descripción

Este proyecto es un bot para Discord modular y de alto rendimiento construido en **Go**. Aprovecha el sistema nativo de **Slash Commands (`/`)** e interacciones de Discord para ofrecer una experiencia fluida, segura y con baja latencia para comunidades y servidores.

---

## ✨ Características Actuales

- [x] **Slash Commands (`/`) Nativos:** Integración completa con la API de interacciones de Discord.
- [x] **Arquitectura Concurrente:** Procesamiento ultrarrápido y bajo consumo de memoria gracias al runtime de Go.
- [x] **Gestión Segura de Configuración:** Carga de credenciales mediante variables de entorno (`.env`).
- [x] **Limpieza Automática:** Registro dinámico y remoción limpia de comandos al detener el proceso.
- [x] **Comandos Básicos:**
  - `/ping`: Mide y muestra la latencia en tiempo real con la pasarela (*Gateway*) de Discord.
  - `/saludar`: Interacción personalizada con mención directa al usuario.

---

## 🗺️ Roadmap / Futuras Implementaciones

El bot se encuentra en desarrollo activo. A continuación se detallan los módulos planificados:

### 🎵 1. Reproducción de Música
- [ ] Reproducción de audio en canales de voz desde múltiples fuentes (YouTube, Spotify, SoundCloud).
- [ ] Control de cola de reproducción (`/play`, `/skip`, `/stop`, `/queue`, `/pause`).
- [ ] Ajuste de volumen y panel de control interactivo mediante botones y selectores.

### 🛡️ 2. Moderación Automática y Control de Lenguaje
- [ ] Filtro inteligente de palabras prohibidas y enlaces no permitidos.
- [ ] Detección y sanción automática de spam, flood y menciones masivas.
- [ ] Sistema de advertencias (`/warn`), silencios temporales (`/timeout`) y baneos (`/ban`).
- [ ] Registro de auditoría (*Audit Logs*) en canales dedicados de moderación.

### 📩 3. Mensajería Avanzada y Utilidades de Comunidad
- [ ] **Mensajes de bienvenida y despedida:** Banners dinámicos y asignación automática de roles (*Autorole*).
- [ ] **Constructor de Embeds (`/embed`):** Creación de mensajes estructurados con títulos, colores, campos e imágenes.
- [ ] **Sistema de Anuncios y Notificaciones:** Envío programado y difusión de mensajes a canales específicos.
- [ ] **Tickets de Soporte:** Creación de canales privados de atención al usuario con transcripción automática.

### 🧠 4. Integración con Inteligencia Artificial (IA)
- [ ] Conexión a APIs de LLMs (OpenAI / Gemini / Anthropic) para chat contextual y respuestas inteligentes.
- [ ] Resumen automático de conversaciones largas en canales de texto.
- [ ] Asistente de consultas y generación de contenido directo en el chat.

### 🔍 5. Búsqueda y Herramientas Web
- [ ] Búsqueda directa en Google / DuckDuckGo dentro de Discord (`/search`).
- [ ] Búsqueda y previsualización de imágenes, GIFs y contenido multimedia.
- [ ] Consulta de clima, conversor de divisas y utilidades de red en tiempo real.

---

## 🛠️ Tecnologías Utilizadas

| Componente | Tecnología |
| :--- | :--- |
| **Lenguaje principal** | [Go (Golang)](https://go.dev/) |
| **Librería de Discord** | [Disgo](https://github.com/disgoorg/disgo@latest) |
| **Gestor de Entorno** | [Godotenv](https://github.com/joho/godotenv) |
| **Arquitectura** | Event-Driven & Interaction Handlers |

---

## 🚀 Instalación y Despliegue

### Prerrequisitos

- **Go** (versión 1.21 o superior).
- Una aplicación y bot creados en el [Discord Developer Portal](https://discord.com/developers/applications).

### Pasos

1. **Clonar el repositorio:**
   ```bash
   git clone [https://github.com/tu-usuario/tu-repositorio.git](https://github.com/tu-usuario/tu-repositorio.git)
   cd tu-repositorio
   ```

2. **Instalar dependencias:**
   ```bash
   go mod download
   ```

3. **Configurar el archivo `.env`:**  
   Crea un archivo `.env` en la raíz del proyecto basándote en el ejemplo:
   ```env
   TOKEN_BOT=tu_token_secreto_aqui
   ```

4. **Ejecutar el bot:**
   ```bash
   go run main.go
   ```

---

## ⚙️ Configuración del Bot en Discord

Para que los comandos de barra (`/`) funcionen correctamente:
1. Dirígete a **OAuth2 $\rightarrow$ URL Generator** en el portal de desarrolladores.
2. En **Scopes**, marca:
   - `bot`
   - `applications.commands`
3. En **Bot Permissions**, selecciona los permisos adecuados (ej. *Send Messages*, *Embed Links*, *Connect*, *Speak*).
4. Usa la URL generada para invitar el bot a tu servidor.

---

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.
