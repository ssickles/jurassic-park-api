#!/bin/sh

echo "Trying to add a cage with an invalid power status"
curl -X POST -H "Content-Type: application/json" -d '{"name":"east-1", "capacity":2, "powerStatus": "INVALID"}' http://localhost:8888/cages
echo ""
echo "Adding cage east-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"east-1", "capacity":2, "powerStatus": "ACTIVE"}' http://localhost:8888/cages
echo ""
echo "Adding cage west-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"west-1", "capacity":10}' http://localhost:8888/cages
echo ""
echo "Adding cage north-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"north-1", "capacity":10, "powerStatus": "ACTIVE"}' http://localhost:8888/cages
echo ""
echo "Adding dinosaur Val, no cage assignment"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Val", "speciesName":"Velociraptor"}' http://localhost:8888/dinosaurs
echo ""
echo "Adding dinosaur Vincent in cage east-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Vincent", "speciesName":"Velociraptor", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
echo "Trying to add dinosaur Steve in cage east-1 but food type is carnivore"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Steve", "speciesName":"Stegosaurus", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
echo "Adding dinosaur Steve in cage west-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Steve", "speciesName":"Stegosaurus", "cageName":"west-1"}' http://localhost:8888/dinosaurs
echo ""
echo "Trying to add dinosaur Steve but already exists"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Steve", "speciesName":"Stegosaurus", "cageName":"west-1"}' http://localhost:8888/dinosaurs
echo ""
echo "Trying to add dinosaur Matt but species does not exist"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Matt", "speciesName":"Mega"}' http://localhost:8888/dinosaurs
echo ""
echo "Adding dinosaur Matt in cage east-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Matt", "speciesName":"Megalosaurus", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
echo "Adding dinosaur Sandy in cage west-1"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Sandy", "speciesName":"Stegosaurus", "cageName":"west-1"}' http://localhost:8888/dinosaurs
echo ""
echo "Trying to add dinosaur Scott in cage east-1 but cage is full"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Scott", "speciesName":"Spinosaurus", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
echo "View dinosaurs for cage east-1"
curl http://localhost:8888/cages/1/assignments
echo ""
echo "Move dinosaur Sandy to cage north-1"
curl -X POST -H "Content-Type: application/json" -d '{"dinosaurName":"Sandy"}' http://localhost:8888/cages/3/assignments
echo ""
echo "View dinosaurs for cage east-1"
curl http://localhost:8888/cages/1/assignments
echo ""
echo "View dinosaurs for cage west-1"
curl http://localhost:8888/cages/2/assignments
echo ""
echo "View dinosaurs for cage north-1"
curl http://localhost:8888/cages/3/assignments
echo ""
echo "View all cages"
curl http://localhost:8888/cages
echo ""
echo "View all dinosaurs"
curl http://localhost:8888/dinosaurs
echo ""