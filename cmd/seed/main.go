package main

import (
	"context"
	"fmt"
	"os"

	"example.com/config"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	godotenv.Load()

	ctx := context.Background()
	url := os.Getenv("NEO4J_URL")
	user := os.Getenv("NEO4J_USER")
	pass := os.Getenv("NEO4J_PASS")

	driver := config.NewDB(url, user, pass)
	err := driver.Connect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established")

	result, err := neo4j.ExecuteQuery(ctx, driver.GetDB(),
		`
		CREATE (n5:Estacion {nombre: "San Antonio"})-[:CONECTA_CON {tiempo: 2}]->(n4:Estacion {nombre: "Las Minas"})-[:CONECTA_CON {tiempo: 2}]->(n5),
		(n3:Estacion {nombre: "Carrizal"})-[:CONECTA_CON {tiempo: 2}]->(n4)-[:CONECTA_CON {tiempo: 2}]->(n3)-[:CONECTA_CON {tiempo: 2}]->(n2:Estacion {nombre: "Los Cerritos"})-[:CONECTA_CON {tiempo: 2}]->(n3),
		(n2)-[:CONECTA_CON {tiempo: 2}]->(n1:Estacion {nombre: "Casco Central II"})-[:CONECTA_CON {tiempo: 2}]->(n2),
		(n1)-[:CONECTA_CON {tiempo:  2}]->(n0:Estacion {nombre: "Casco Central"})-[:CONECTA_CON {tiempo: 2}]->(n1),
		(n0)-[:CONECTA_CON {tiempo: 2}]->(Estacion:Estacion {nombre: "El Tambor"})-[:CONECTA_CON {tiempo: 2}]->(n0),
		(Estacion)-[:CONECTA_CON {tiempo: 3}]->(n7:Estacion {nombre: "Ruiz Pineda"})-[:CONECTA_CON {tiempo: 3}]->(Estacion),
		(n7)-[:CONECTA_CON {tiempo: 3}]->(n8:Estacion {nombre: "Mamera"})-[:CONECTA_CON {tiempo: 3}]->(n7),
		(n8)-[:CONECTA_CON {tiempo: 2}]->(n9:Estacion {nombre: "Antimano"})-[:CONECTA_CON {tiempo: 2}]->(n8)<-[:CONECTA_CON {tiempo: 3}]-(n35:Estacion {nombre: "Caricuao"})<-[:CONECTA_CON {tiempo: 3}]-(n8),
		(n9)-[:CONECTA_CON {tiempo: 3}]->(n10:Estacion {nombre: "Carapita"})-[:CONECTA_CON {tiempo: 3}]->(n9),
		(n10)-[:CONECTA_CON {tiempo: 3}]->(n11:Estacion {nombre: "La Yaguara"})-[:CONECTA_CON {tiempo: 3}]->(n10),
		(n11)-[:CONECTA_CON {tiempo: 3}]->(n12:Estacion {nombre: "La Paz"})-[:CONECTA_CON {tiempo: 3}]->(n11),
		(n12)-[:CONECTA_CON {tiempo: 2}]->(n13:Estacion {nombre: "Artigas"})-[:CONECTA_CON {tiempo: 2}]->(n12),
		(n13)-[:CONECTA_CON {tiempo: 2}]->(n14:Estacion {nombre: "Maternidad"})-[:CONECTA_CON {tiempo: 2}]->(n13),
		(n14)-[:CONECTA_CON {tiempo: 2}]->(n15:Estacion {nombre: "Capuchinos"})-[:CONECTA_CON {tiempo: 2}]->(n14),
		(n15)-[:CONECTA_CON {tiempo: 2}]->(n16:Estacion {nombre: "Teatros"})-[:CONECTA_CON {tiempo: 2}]->(n15)-[:CONECTA_CON {tiempo: 2}]->(n51:Estacion {nombre: "Capitolio"})-[:CONECTA_CON {tiempo: 2}]->(n15),
		(n18:Estacion {nombre: "Bello Monte"})-[:CONECTA_CON {tiempo: 2}]->(n19:Estacion {nombre: "Las Mercedes"})-[:CONECTA_CON {tiempo: 2}]->(n18)<-[:CONECTA_CON {tiempo: 4}]-(n88:Estacion {nombre: "Zona Rental"})<-[:CONECTA_CON {tiempo: 4}]-(n18),
		(n19)-[:CONECTA_CON {tiempo: 2}]->(n20:Estacion {nombre: "Tamanaco"})-[:CONECTA_CON {tiempo: 2}]->(n19),
		(n20)-[:CONECTA_CON {tiempo: 2}]->(n21:Estacion {nombre: "Chuao"})-[:CONECTA_CON {tiempo: 2}]->(n20)-[:CONECTA_CON {tiempo: 2}]->(n36:Estacion {nombre: "San Roman"})-[:CONECTA_CON {tiempo: 2}]->(n20),
		(n21)-[:CONECTA_CON {tiempo: 4}]->(n22:Estacion {nombre: "Bello Campo"})-[:CONECTA_CON {tiempo: 4}]->(n21)-[:CONECTA_CON {tiempo: 2}]->(n46:Estacion {nombre: "Santa Marta"})-[:CONECTA_CON {tiempo: 2}]->(n21),
		(n35)-[:CONECTA_CON {tiempo: 13}]->(n34:Estacion {nombre: "Terminal Caracas"})-[:CONECTA_CON {tiempo: 13}]->(n35),
		(n34)-[:CONECTA_CON {tiempo: 4}]->(n33:Estacion {nombre: "Mercado"})-[:CONECTA_CON {tiempo: 4}]->(n34),
		(n33)-[:CONECTA_CON {tiempo: 2}]->(n32:Estacion {nombre: "Coche"})-[:CONECTA_CON {tiempo: 2}]->(n33),
		(n32)-[:CONECTA_CON {tiempo: 3}]->(n31:Estacion {nombre: "Jardines"})-[:CONECTA_CON {tiempo: 3}]->(n32),
		(n31)-[:CONECTA_CON {tiempo: 4}]->(n27:Estacion {nombre: "El Valle"})-[:CONECTA_CON {tiempo: 4}]->(n31),
		(n27)-[:CONECTA_CON {tiempo: 2}]->(n28:Estacion {nombre: "La Bandera"})-[:CONECTA_CON {tiempo: 2}]->(n27)-[:CONECTA_CON {tiempo: 2}]->(n26:Estacion {nombre: "Prado De Maria"})-[:CONECTA_CON {tiempo: 2}]->(n27),
		(n28)-[:CONECTA_CON {tiempo: 2}]->(n29:Estacion {nombre: "Los Simbolos"})-[:CONECTA_CON {tiempo: 2}]->(n28),
		(n29)-[:CONECTA_CON {tiempo: 2}]->(n30:Estacion {nombre: "Ciudad Universitaria"})-[:CONECTA_CON {tiempo: 2}]->(n29),
		(n26)-[:CONECTA_CON {tiempo: 2}]->(n25:Estacion {nombre: "El Peaje"})-[:CONECTA_CON {tiempo: 2}]->(n26),
		(n25)-[:CONECTA_CON {tiempo: 2}]->(n23:Estacion {nombre: "San Agustin"})-[:CONECTA_CON {tiempo: 2}]->(n25),
		(n36)-[:CONECTA_CON {tiempo: 2}]->(n37:Estacion {nombre: "Santa Fe"})-[:CONECTA_CON {tiempo: 2}]->(n36)-[:CONECTA_CON {tiempo: 7}]->(n46)-[:CONECTA_CON {tiempo: 7}]->(n36),
		(n37)-[:CONECTA_CON {tiempo: 2}]->(n38:Estacion {nombre: "La Ciudadela"})-[:CONECTA_CON {tiempo: 2}]->(n37),
		(n38)-[:CONECTA_CON {tiempo: 2}]->(n39:Estacion {nombre: "Prados Del Este"})-[:CONECTA_CON {tiempo: 2}]->(n38),
		(n39)-[:CONECTA_CON {tiempo: 2}]->(n40:Estacion {nombre: "Iglesia"})-[:CONECTA_CON {tiempo: 2}]->(n39),
		(n41:Estacion {nombre: "Baruta"})-[:CONECTA_CON {tiempo: 3}]->(n40)-[:CONECTA_CON {tiempo: 3}]->(n41)-[:CONECTA_CON {tiempo: 3}]->(n42:Estacion {nombre: "La Trinidad"})-[:CONECTA_CON {tiempo: 3}]->(n41),
		(n42)-[:CONECTA_CON {tiempo: 2}]->(n43:Estacion {nombre: "La Boyera"})-[:CONECTA_CON {tiempo: 2}]->(n42),
		(n43)-[:CONECTA_CON {tiempo: 2}]->(n44:Estacion {nombre: "Los Geranios"})-[:CONECTA_CON {tiempo: 2}]->(n43),
		(n44)-[:CONECTA_CON {tiempo: 2}]->(:Estacion {nombre: "El Hatillo"})-[:CONECTA_CON {tiempo: 2}]->(n44),
		(n46)-[:CONECTA_CON {tiempo: 2}]->(n47:Estacion {nombre: "San Luis"})-[:CONECTA_CON {tiempo: 2}]->(n46),
		(n47)-[:CONECTA_CON {tiempo: 2}]->(n48:Estacion {nombre: "Santa Paula"})-[:CONECTA_CON {tiempo: 2}]->(n47),
		(n48)-[:CONECTA_CON {tiempo: 2}]->(:Estacion {nombre: "Santa Ana"})-[:CONECTA_CON {tiempo: 2}]->(n48),
		(n55:Estacion {nombre: "Perez Bonalde"})-[:CONECTA_CON {tiempo: 2}]->(n54:Estacion {nombre: "Plaza Sucre"})-[:CONECTA_CON {tiempo: 2}]->(n55),
		(n53:Estacion {nombre: "Agua Salud"})-[:CONECTA_CON {tiempo: 2}]->(n52:Estacion {nombre: "Caño Amarillo"})-[:CONECTA_CON {tiempo: 2}]->(n53)<-[:CONECTA_CON {tiempo: 2}]-(n85:Estacion {nombre: "Gato Negro"})<-[:CONECTA_CON {tiempo: 2}]-(n53)-[:CONECTA_CON {tiempo: 3}]->(n67:Estacion {nombre: "Miraflores"})-[:CONECTA_CON {tiempo: 3}]->(n53),
		(n52)-[:CONECTA_CON {tiempo: 2}]->(n51)-[:CONECTA_CON {tiempo: 2}]->(n52),
		(n56:Estacion {nombre: "Parque Carabobo"})-[:CONECTA_CON {tiempo: 2}]->(n57:Estacion {nombre: "Bellas Artes"})-[:CONECTA_CON {tiempo: 2}]->(n56)<-[:CONECTA_CON {tiempo: 2}]-(n86:Estacion {nombre: "La Hoyada"})<-[:CONECTA_CON {tiempo: 2}]-(n56),
		(n57)-[:CONECTA_CON {tiempo: 2}]->(n58:Estacion {nombre: "Colegio De Ingenieros"})-[:CONECTA_CON {tiempo: 2}]->(n57),
		(n59:Estacion {nombre: "Sabana Grande"})-[:CONECTA_CON {tiempo: 2}]->(n60:Estacion {nombre: "Chacaito"})-[:CONECTA_CON {tiempo: 2}]->(n59)<-[:CONECTA_CON {tiempo: 3}]-(n88)<-[:CONECTA_CON {tiempo: 3}]-(n59),
		(n60)-[:CONECTA_CON {tiempo: 3}]->(n61:Estacion {nombre: "Chacao"})-[:CONECTA_CON {tiempo: 3}]->(n60),
		(n61)-[:CONECTA_CON {tiempo: 2}]->(n62:Estacion {nombre: "Altamira"})-[:CONECTA_CON {tiempo: 2}]->(n61),
		(n63:Estacion {nombre: "Los Cortijos"})-[:CONECTA_CON {tiempo: 3}]->(n64:Estacion {nombre: "La California"})-[:CONECTA_CON {tiempo: 3}]->(n63)-[:CONECTA_CON {tiempo: 2}]->(n90:Estacion {nombre: "Los Dos Caminos"})-[:CONECTA_CON {tiempo: 2}]->(n63),
		(n64)-[:CONECTA_CON {tiempo: 3}]->(:Estacion {nombre: "Petare"})-[:CONECTA_CON {tiempo: 3}]->(n64),
		(n67)-[:CONECTA_CON {tiempo: 3}]->(n24:Estacion {nombre: "Urdaneta"})-[:CONECTA_CON {tiempo: 3}]->(n67),
		(n24)-[:CONECTA_CON {tiempo: 2}]->(n68:Estacion {nombre: "San Bernardino"})-[:CONECTA_CON {tiempo: 2}]->(n24)-[:CONECTA_CON {tiempo: 2}]->(n86)-[:CONECTA_CON {tiempo: 2}]->(n24),
		(n68)-[:CONECTA_CON {tiempo: 2}]->(n69:Estacion {nombre: "Guaicaipuro"})-[:CONECTA_CON {tiempo: 2}]->(n68),
		(n69)-[:CONECTA_CON {tiempo: 2}]->(n70:Estacion {nombre: "Mariperez"})-[:CONECTA_CON {tiempo: 2}]->(n69),
		(n70)-[:CONECTA_CON {tiempo: 2}]->(n71:Estacion {nombre: "Las Palmas"})-[:CONECTA_CON {tiempo: 2}]->(n70),
		(n72:Estacion {nombre: "Chapellin"})-[:CONECTA_CON {tiempo: 2}]->(n73:Estacion {nombre: "La Castellana"})-[:CONECTA_CON {tiempo: 2}]->(n72)<-[:CONECTA_CON {tiempo: 2}]-(n71)<-[:CONECTA_CON {tiempo: 2}]-(n72),
		(n73)-[:CONECTA_CON {tiempo: 2}]->(n74:Estacion {nombre: "Avila"})-[:CONECTA_CON {tiempo: 2}]->(n73),
		(n75:Estacion {nombre: "Santa Eduvigis"})-[:CONECTA_CON {tiempo: 2}]->(n74)-[:CONECTA_CON {tiempo: 2}]->(n75)-[:CONECTA_CON {tiempo: 2}]->(n89:Estacion {nombre: "Montecristo"})-[:CONECTA_CON {tiempo: 2}]->(n75),
		(n76:Estacion {nombre: "Boleita"})-[:CONECTA_CON {tiempo: 2}]->(n77:Estacion {nombre: "Horizonte"})-[:CONECTA_CON {tiempo: 2}]->(n76)-[:CONECTA_CON {tiempo: 2}]->(n90)-[:CONECTA_CON {tiempo: 2}]->(n76),
		(n77)-[:CONECTA_CON {tiempo: 2}]->(n78:Estacion {nombre: "La Urbina Sur"})-[:CONECTA_CON {tiempo: 2}]->(n77),
		(n78)-[:CONECTA_CON {tiempo: 2}]->(n79:Estacion {nombre: "La Urbina"})-[:CONECTA_CON {tiempo: 2}]->(n78),
		(n79)-[:CONECTA_CON {tiempo: 2}]->(n80:Estacion {nombre: "Tamarindo"})-[:CONECTA_CON {tiempo: 2}]->(n81:Estacion {nombre: "Guarenas I"})-[:CONECTA_CON {tiempo: 2}]->(n82:Estacion {nombre: "Guarenas II"})-[:CONECTA_CON {tiempo: 2}]->(n83:Estacion {nombre: "Guatire I"})-[:CONECTA_CON {tiempo: 2}]->(:Estacion {nombre: "Guatire II"})-[:CONECTA_CON {tiempo: 2}]->(n83)-[:CONECTA_CON {tiempo: 2}]->(n82)-[:CONECTA_CON {tiempo: 2}]->(n81)-[:CONECTA_CON {tiempo: 2}]->(n80)-[:CONECTA_CON {tiempo: 2}]->(n79),
		(n54)-[:CONECTA_CON {tiempo: 2}]->(n85)-[:CONECTA_CON {tiempo: 2}]->(n54),
		(n66:Estacion {nombre: "Los Magallanes"})-[:CONECTA_CON {tiempo: 3}]->(n85)-[:CONECTA_CON {tiempo: 3}]->(n66),
		(n51)-[:CONECTA_CON {tiempo: 2}]->(n86)-[:CONECTA_CON {tiempo: 2}]->(n51),
		(n87:Estacion {nombre: "Nuevo Circo"})-[:CONECTA_CON {tiempo: 2}]->(n17:Estacion {nombre: "Parque Central"})-[:CONECTA_CON {tiempo: 2}]->(n87)-[:CONECTA_CON {tiempo: 2}]->(n16)-[:CONECTA_CON {tiempo: 2}]->(n87),
		(n86)-[:CONECTA_CON {tiempo: 2}]->(n17)-[:CONECTA_CON {tiempo: 2}]->(n86)<-[:CONECTA_CON {tiempo: 2}]-(n23)<-[:CONECTA_CON {tiempo: 2}]-(n86),
		(n17)-[:CONECTA_CON {tiempo: 3}]->(n88)-[:CONECTA_CON {tiempo: 3}]->(n17),
		(n58)-[:CONECTA_CON {tiempo: 2}]->(n88)-[:CONECTA_CON {tiempo: 2}]->(n58),
		(n30)-[:CONECTA_CON {tiempo: 3}]->(n88)-[:CONECTA_CON {tiempo: 3}]->(n30),
		(n62)-[:CONECTA_CON {tiempo: 2}]->(n89)-[:CONECTA_CON {tiempo: 2}]->(n62),
		(n89)-[:CONECTA_CON {tiempo: 2}]->(n90)-[:CONECTA_CON {tiempo: 2}]->(n89)<-[:CONECTA_CON {tiempo: 3}]-(n22)<-[:CONECTA_CON {tiempo: 3}]-(n89)
	`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		panic(err)
	}

	_ = result

	fmt.Println("Done Seeding")

	defer driver.Close(ctx)

}
