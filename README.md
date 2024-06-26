# Algoritmos y Estructura de Datos - cátedra Buchwald (75.41)

- Los lenguajes utilizados fueron GO y Python
- La cursada consta de un total de 11 entregas obligatorias que se listan debajo
> Si estas cursando la materia en esta cátedra por primera vez y aun no entraste en crisis un sabado 4am por una entrega que no pasa el corrector automatico, recomiendo fuertemente seguir intentando que lo vas a resolver en el momento mas esquizofrenico y siempre podes preguntarle a tu ayudante o en el channel general de slack

## Entregas 2022C2
- [TP0](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/TP0)
- [Administración de Memoria](https://github.com/Igris-1/Algoritmos-y-estructura-de-datos/tree/main/adm)
- [Pila](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/pila)
- [Cola](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/cola)
- [Lista](lista)
- [TP1](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/TP1)
- [Hash](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/hash)
- [ABB](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/abb)
- [Heap](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/heap)
- [TP2](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/TP2)
- [TP3](https://github.com/Igris-1/Algoritmos-y-programacion-II/tree/main/TP3)

### Para correr TP1, 2 y 3, estando dentro del modulo del programa

- TP1
```
$ ./rerepolez lista_candidatos.csv padron.csv
<Se queda esperando por comandos>
```

- TP2
```
$ ./algogram usuarios.txt
<Se queda esperando por comandos>
```
- TP3
```
$ ./vamosmoshi ciudades.pj
<Se queda esperando por comandos>
```

> Correr pruebas TP1 - TP2
```
$ ./pruebas.sh PATH-A-EJECUTABLE-TP1-TP2
```

### Test Hash y ABB
- Para ejecutar las pruebas, incluyendo las pruebas de volumen (benchmarks, que toman los tiempos y consumos de memoria), ejecutar:
```
$ go test -bench=. -benchmem
```
