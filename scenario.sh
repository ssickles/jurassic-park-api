#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{"name":"east-1", "capacity":2, "powerStatus": "INVALID"}' http://localhost:8888/cages
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"east-1", "capacity":2, "powerStatus": "ACTIVE"}' http://localhost:8888/cages
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"west-1", "capacity":10}' http://localhost:8888/cages
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Val", "speciesName":"Velociraptor"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Vincent", "speciesName":"Velociraptor", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Steve", "speciesName":"Stegosaurus", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Steve", "speciesName":"Stegosaurus", "cageName":"west-1"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Steve", "speciesName":"Stegosaurus", "cageName":"west-1"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Matt", "speciesName":"Mega"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Matt", "speciesName":"Megalosaurus", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Sandy", "speciesName":"Stegosaurus", "cageName":"west-1"}' http://localhost:8888/dinosaurs
echo ""
curl -X POST -H "Content-Type: application/json" -d '{"name":"Scott", "speciesName":"Spinosaurus", "cageName":"east-1"}' http://localhost:8888/dinosaurs
echo ""
