# Recruitment task Goonline - Go Gin Rest API.

### How to run the project?
First of all you need to have a MySQL instance running on your PC on ```port 3306```, with a schema named ```goonline```.

Then simply run ```go run main.go```.

The app will automatically populate the database with Pokemons from provided json file and run the server on the ```port 8080``` of your localhost.

Public route (for an unauthenticated user) is named ```/public```, while all others are under the ```/public/api``` route.

