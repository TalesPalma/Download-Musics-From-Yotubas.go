# API para Download de Músicas do YouTube

Esta é uma API em Go que permite baixar músicas do YouTube. Siga as instruções abaixo para configurar e usar a API.

## Pré-requisitos

- Go 1.16 ou superior
- Uma conta no YouTube
- ffmpeg https://www.ffmpeg.org/download.html

## Instalação

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/youtube-music-downloader-api.git
   ```

2. Entre no diretório do projeto:
   ```
   cd youtube-music-downloader-api
   ```

3. Instale as dependências:
   ```
   go mod tidy
   ```

## Configuração

## Uso

Para usar a API, siga os passos abaixo:

1. Inicie o servidor:
   ```
   go run main.go
   ```

2. A API estará disponível em `http://localhost:8080` por padrão.

3. Para baixar músicas de uma playlist do YouTube, faça uma requisição POST para o endpoint `/download` com o seguinte corpo JSON:

   ```json
   {
     "playlist_url": "https://www.youtube.com/playlist?list=SEU_ID_DA_PLAYLIST"
   }
   ```

   Substitua `SEU_ID_DA_PLAYLIST` pelo ID real da playlist que você deseja baixar.

4. A API irá processar a solicitação e baixar as músicas da playlist especificada.

5. As músicas baixadas serão salvas no diretório `downloads` dentro do projeto.

Nota: Certifique-se de que você tem permissão para baixar o conteúdo da playlist e que está em conformidade com os termos de serviço do YouTube.


## Contribuição

Se você encontrar algum problema ou tiver sugestões para melhorar a API, sinta-se à vontade para abrir uma issue ou enviar um pull request.
