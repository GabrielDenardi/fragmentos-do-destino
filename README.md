# 🌌 Fragmentos do Destino

**Fragmentos do Destino** é um projeto minimalista que une filosofia, computação e o universo para proporcionar uma experiência imersiva e reflexiva. A aplicação exibe fragmentos filosóficos e perguntas retóricas que se transformam dinamicamente conforme a interação do usuário, complementadas por um ambiente visual e sonoro inspirado na vastidão do cosmos.

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)  
[![Feito com Go](https://img.shields.io/badge/backend-Go-informational?logo=go)](https://golang.org)  
[![Frontend](https://img.shields.io/badge/frontend-HTML%2FCSS%2FJS-orange)](#)

---

## ✨ Funcionalidades

- 🌠 **Fundo Estelar Dinâmico**  
  Animação de estrelas em um canvas com fundo que se adapta ao horário (mais claro durante o dia e mais escuro à noite).

- 🧠 **Fragmentos Filosóficos e Perguntas Retóricas**  
  Exibe citações e provocações reflexivas. Após um número definido de interações, ativa o modo retórico, que gera novas perguntas dinâmicas — por lista interna ou via API da OpenAI (se configurada).

- 🫱 **Interação do Usuário**  
  Botões **"Aceito"** e **"Rejeito"** registram as escolhas do usuário e atualizam os fragmentos apresentados.

- 📤 **Compartilhamento Inspirador**  
  Gera uma imagem do fragmento atual para ser compartilhada nas redes sociais.

- 🎧 **Ambiente Sonoro Imersivo**  
  Sons de ambientação e sussurros filosóficos amplificam a imersão sensorial.

- 🔄 **Modo Retórico Dinâmico**  
  Ativado após certo número de cliques, cria perguntas inéditas com base em listas locais ou por meio da IA.

---

## 🧩 Estrutura do Projeto

```plaintext
fragmentos-do-destino/
├── backend/
│   ├── main.go            # Código do servidor Go
│   ├── go.mod             # Gerenciamento de dependências
│   ├── go.sum             # Sumário das dependências
│   └── .env               # Variáveis de ambiente (ex: OPENAI_API_KEY)
├── frontend/
│   ├── index.html         # Página principal da aplicação
│   ├── 001210_the-sound-of-universe-53546.mp3  # Áudio de ambientação
│   ├── superspacy-atmosphere-106826.mp3        # Áudio do sussurro filosófico
└── README.md              # Este arquivo
```

---

## ⚙️ Como Executar

### Pré-requisitos

- Go  
- Navegador moderno  
- (Opcional) Chave da API da OpenAI  

### Passos

Clone o repositório:

```bash
git clone https://github.com/GabrielDenardi/fragmentos-do-destino.git
cd fragmentos-do-destino
```

Configure o arquivo `.env` no diretório `backend/` com sua chave (caso use a API da OpenAI):

```env
OPENAI_API_KEY=sua_chave_aqui
```

Execute o servidor Go:

```bash
cd backend
go run main.go
```

Abra o arquivo `frontend/index.html` no navegador ou sirva o diretório com um servidor estático.

---

## 🌌 Inspiração

Este projeto nasceu do desejo de transcender o utilitário da tecnologia e criar uma ponte entre o humano e o cósmico. Um espaço para contemplar, aceitar, rejeitar — e, talvez, se transformar através das palavras.

---

## 📜 Licença

Distribuído sob a Licença MIT. Veja o arquivo [`LICENSE`](LICENSE) para mais informações.

---

## 🤝 Contribuições

Sinta-se livre para contribuir com melhorias, ideias de fragmentos filosóficos, novos efeitos visuais ou integrações com IA. Toda ajuda é bem-vinda!

> *"Aceitar ou rejeitar não é apenas um gesto, mas um movimento da alma."*
