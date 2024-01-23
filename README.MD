docker compose up -d
docker system prune
docker system prune --all --force --volumes


# Migrate 
#### Adicionar a chave GPG para a biblioteca golang-migrate
    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | gpg --dearmor -o /usr/share/keyrings/migrate-archive-keyring.gpg

#### Adicionar o repositório golang-migrate ao sistema de pacotes
    echo "deb [signed-by=/usr/share/keyrings/migrate-archive-keyring.gpg] https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/migrate.list > /dev/null

#### Atualizar o índice de pacotes
    sudo apt-get update
    
    # Instalar a ferramenta migrate
    sudo apt-get install -y migrate
    
#### Verificar a versão da ferramenta migrate
    migrate -version

#### Criar um novo arquivo de migração SQL no diretório db/migrations
    migrate create -ext sql -dir db/migration -seq initial_tables

#### Colocar o código sql de criação da tabela nos arquivos de migração gerados com initial_tables
    000001_initial_tables.down.sql
    000001_initial_tables.up.sql    

#### Aplicar as migrações para o banco de dados MySQL
    migrate -path db/migration -database "mysql://user:pass@tcp(localhost:3306)/social_network" -verbose up

# Fim migrate

# Gorilla mux

#### Instalação
    go get -u github.com/gorilla/mux
# Fim Gorilla mux

# MySQL
#### Acessar o mySQL dentro do container
    mysql -u user -p
    show databases;
    USE social_network;
    desc users;
# Fim mySQL