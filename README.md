# Fragmentos do Destino

**Fragmentos do Destino** é um projeto minimalista que une filosofia, computação e o universo para proporcionar uma experiência imersiva e reflexiva. A aplicação exibe fragmentos filosóficos e perguntas retóricas que se transformam dinamicamente conforme a interação do usuário, complementadas por um ambiente visual e sonoro inspirado na vastidão do cosmos.

## Funcionalidades

- **Fundo Estelar Dinâmico:**  
  Animação de estrelas em um canvas com fundo que se adapta ao horário (mais claro durante o dia e mais escuro à noite).

- **Fragmentos Filosóficos:**  
  Exibe citações e perguntas retóricas. Após um número definido de interações, ativa o modo retórico, gerando novas perguntas retóricas de forma dinâmica – inicialmente por meio de uma lista e, quando configurado, utilizando integração com a API do OpenAI.

- **Interação do Usuário:**  
  Botões "Aceito" e "Rejeito" que registram as escolhas do usuário e atualizam os fragmentos exibidos.

- **Compartilhamento:**  
  Gera uma imagem do fragmento atual para que o usuário possa compartilhar a mensagem em redes sociais.

- **Ambiente Sonoro:**  
  Sons de ambientação e sussurros filosóficos para intensificar a experiência sensorial.

- **Modo Retórico Dinâmico:**  
  Ao atingir um certo número de cliques, o sistema ativa o modo retórico e gera perguntas inéditas – inicialmente de forma dinâmica e, se configurado, utilizando uma API de IA para que as perguntas nunca sejam as mesmas.

## Estrutura do Projeto

```plaintext
fragmentos-do-destino/
├── backend/
│   ├── main.go          # Código do servidor Go
│   ├── go.mod           # Gerenciamento de dependências
│   ├── go.sum           # Sumário das dependências
│   └── .env             # Variáveis de ambiente (ex: OPENAI_API_KEY)
├── frontend/
│   ├── index.html       # Página principal da aplicação
│   ├── 001210_the-sound-of-universe-53546.mp3  # Áudio de ambientação
│   ├── superspacy-atmosphere-106826.mp3         # Áudio do sussurro filosófico
│└── README.md            # Este arquivo
