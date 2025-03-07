para correr este programa es necesario tener instalado en la terminal el paquete de GO,
personalmente utilizo WSL, 

- llamar en la terminal el engine para poder realizar los llamados al cliente
    -- go get -u github.com/gin-gintonic/gin

- tambien instale el sqlite3 para crear una base de datos
    -- go get -u github.com/driver/sqlite3


los agrege go.mod como required

-- realizo un "go mod tidy"(en la terminal) para reorganizar las dependencias y librerias

--> Para correr "go run ." o queres realizar un ejecutable "go build ."

una vez que este corriendo, abrir localhost:8080 en cualquier browser y probar, faltan mejorar cosas.