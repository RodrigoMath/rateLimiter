Rate Limiter em Go com Redis
Este projeto implementa um sistema de Rate Limiting robusto utilizando a linguagem Go, o banco de dados Redis e seguindo os princípios da Clean Architecture. O sistema é capaz de limitar requisições baseadas no endereço IP do cliente ou em um Token de acesso customizado enviado via Header.

Tecnologias Utilizadas
Go (Golang) 1.24+

Redis (Persistência e contagem atômica)

Docker e Docker Compose (Orquestração de ambiente)

Viper (Gerenciamento de configurações e .env)

Arquitetura do Projeto
O projeto foi desenhado para ser extensível e fácil de manter, separando as preocupações em camadas:

Config: Centraliza o carregamento do arquivo .env de forma tipada.

Infra: Implementação do repositório Redis e Factory de banco de dados.

UseCase: Regra de negócio central (lógica de contagem e decisão de bloqueio).

Middleware: Interceptador HTTP que valida o limite antes de processar a requisição.

Como Rodar o Projeto
1. Pré-requisitos
Ter o Docker e o Docker Compose instalados.

2. Configuração do Ambiente
Certifique-se de que o arquivo .env na raiz do projeto possui as definições básicas. Exemplo:
STRATEGY=redis
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=root
LIMIT_IP=10
LIMIT_TOKEN=100
BLOCK_TIME=60

3. Subindo a Aplicação
Execute o comando abaixo para compilar a aplicação e subir o Redis simultaneamente:

docker-compose up --build

O servidor estará disponível em: http://localhost:8080

Como Testar
Para validar o funcionamento, utilize os comandos abaixo em um terminal Linux, WSL ou Git Bash:

Teste por IP (Limite de 10 requisições)
Este comando faz 12 requisições seguidas. O esperado é receber dez códigos 200 e dois códigos 429.

for i in {1..12}; do curl -s -o /dev/null -w "%{http_code} " http://localhost:8080; done

Teste por Token (Limite de 100 requisições)
Este comando envia o header API_KEY. O esperado é receber cem códigos 200 e dez códigos 429, provando que o limite de Token tem precedência.

for i in {1..110}; do curl -s -o /dev/null -w "%{http_code} " -H "API_KEY: meu-token-vip" http://localhost:8080; done

Detalhes das Configurações
LIMIT_IP: Quantidade máxima de requisições por IP.

LIMIT_TOKEN: Quantidade máxima de requisições por Token.

BLOCK_TIME: Tempo em segundos que o usuário ficará bloqueado após exceder o limite.

REDIS_HOST: Endereço do serviço Redis (deve ser "redis" para uso com Docker).

Para resetar os contadores e bloqueios manualmente durante os testes, utilize o comando:
docker exec -it <nome_do_container_redis> redis-cli -a root FLUSHALL