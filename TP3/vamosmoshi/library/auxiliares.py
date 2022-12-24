from tdas import grafo

DESDE = 0
ORIGEN = 0
HASTA = 1


def agregar_ciudades(archivo, grafo, latitudes):
    with open(archivo, "r") as f:
        for fila in f:
            if fila != "":
                elemento = fila.rstrip().split(",")
                if len(elemento) == 3:
                    if not elemento[2].isdigit():
                        latitudes[elemento[0]] = (elemento[1], elemento[2])
                    else:
                        grafo.agregar_vertice(elemento[0])
                        grafo.agregar_vertice(elemento[1])
                        grafo.agregar_arista(elemento[0], elemento[1], int(elemento[2]))


def guardar_kml(ruta, camino, latitudes):
    desde = camino[0]
    hasta = camino[-1]
    with open(ruta, "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        f.write("	<Document>\n")
        if desde == hasta:
            f.write(f"		<name>Camino desde {desde}</name>\n")
        else:
            f.write(f"		<name>Camino desde {desde} hacia {hasta}</name>\n")
        for ciudad in camino:
            longitud, latitud = latitudes[ciudad]
            f.write(f"		<Placemark>\n")
            f.write(f"			<name>{ciudad}</name>\n")
            f.write(f"			<Point>\n")
            f.write(f"				<coordinates>{longitud}, {latitud}</coordinates>\n")
            f.write(f"			</Point>\n")
            f.write(f"		</Placemark>\n")
        for i in range(len(camino) - 1):
            ciudad1 = camino[i]
            ciudad2 = camino[i + 1]
            long1, lat1 = latitudes[ciudad1]
            long2, lat2 = latitudes[ciudad2]
            f.write(f"		<Placemark>\n")
            f.write(f"			<LineString>\n")
            f.write(f"				<coordinates>{long1}, {lat1} {long2}, {lat2}</coordinates>\n")
            f.write(f"			</LineString>\n")
            f.write(f"		</Placemark>\n")
        f.write("	</Document>\n")
        f.write("</kml>\n")



def guardar_viaje_kml(ruta, camino, latitudes):
    desde = camino[0]
    agregados = set()
    aristas = set()
    with open(ruta, "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        f.write("	<Document>\n")
        f.write(f"		<name>Camino desde {desde}</name>\n")

        for ciudad in camino:
            if ciudad not in agregados:
                agregados.add(ciudad)
                longitud, latitud = latitudes[ciudad]
                f.write(f"		<Placemark>\n")
                f.write(f"			<name>{ciudad}</name>\n")
                f.write(f"			<Point>\n")
                f.write(f"				<coordinates>{longitud}, {latitud}</coordinates>\n")
                f.write(f"			</Point>\n")
                f.write(f"		</Placemark>\n")
        
        for i in range(len(camino) - 1):
            ciudad1 = camino[i]
            ciudad2 = camino[i + 1]
            if (ciudad1, ciudad2) not in aristas:
                aristas.add((ciudad1, ciudad2))
                aristas.add((ciudad2, ciudad1))
                long1, lat1 = latitudes[ciudad1]
                long2, lat2 = latitudes[ciudad2]
                f.write(f"		<Placemark>\n")
                f.write(f"			<LineString>\n")
                f.write(f"				<coordinates>{long1}, {lat1} {long2}, {lat2}</coordinates>\n")
                f.write(f"			</LineString>\n")
                f.write(f"		</Placemark>\n")
        f.write("	</Document>\n")
        f.write("</kml>\n")

def recomendaciones(ruta):
    grafito = grafo.Grafo(True)

    with open(ruta, "r") as f:
        for fila in f:
            if fila != "":
                elemento = fila.rstrip().split(",")
                grafito.agregar_vertice(elemento[0])
                grafito.agregar_vertice(elemento[1])
                grafito.agregar_arista(elemento[0], elemento[1])
    return grafito


def guardar_pj(ciudades, ruta, cant, latitudes):
    c_guardar = _guardar_pj(ciudades)
    with open(ruta, "w") as f:
        f.write(f"{len(c_guardar)}\n")
        for ciudad in c_guardar:
            f.write(f"{ciudad},{latitudes[ciudad][0]},{latitudes[ciudad][1]}\n")

        f.write(f"{cant}\n")
        for ciudad in ciudades:
            ciudad1 = ciudad[0]
            ciudad2 = ciudad[1]
            peso = ciudad[2]
            f.write(f"{ciudad1},{ciudad2},{peso}\n")


def _guardar_pj(ciudades):
    c_guardar = []
    for c1, c2, _ in ciudades:
        if c1 not in c_guardar:
            c_guardar.append(c1)
        if c2 not in c_guardar:
            c_guardar.append(c2)
    return c_guardar


def pertenece(ubicacion, grafo):
    if len(ubicacion) == 2:
        if (
            not ubicacion[DESDE] in grafo.obtener_vertices()
            or not ubicacion[HASTA] in grafo.obtener_vertices()
        ):
            return False
    if not ubicacion[ORIGEN] in grafo.obtener_vertices():
        return False
    return True
