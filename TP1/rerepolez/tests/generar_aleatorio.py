import sys
import random

ALTERNATIVAS = ["Presidente", "Gobernador", "Intendente"]
FREQ_IMPUGNAR = 30
FREQ_BLANCO = 50


def ingresar(padron):
    def func():
        return "ingresar " + random.choice(padron)
    return func


def fin_votar():
    return "fin-votar"


def deshacer():
    return "deshacer"


def votar(op_partidos):
    def func():
        impugnar = random.randint(1, FREQ_IMPUGNAR) == 1
        if impugnar:
            return "votar Presidente 0"
        else:
            return "votar Presidente " + str(random.randint(1, op_partidos)) + "\nvotar Gobernador " + str(random.randint(1, op_partidos)) + "\nvotar Intendente " + str(random.randint(1, op_partidos))
    return func


def main(path_padron, cant_opciones, cant_operaciones, salida):
    padron = []
    with open(path_padron) as f:
        for l in f:
            padron.append(l.strip())

    operaciones = [ingresar(padron), ingresar(padron), fin_votar, fin_votar, votar(cant_opciones), votar(cant_opciones), deshacer, ]

    with open(salida, "w") as f:
        for i in range(cant_operaciones):
            se_voto = False
            while not se_voto:
                op = random.choice(operaciones)
                result = op()
                if result.startswith("fin") and random.randint(1, FREQ_BLANCO) != 1:
                    continue
                f.write(result + "\n")
                if result.startswith("votar"):
                    se_voto = True
            if random.randint(1, 5) != 1:
                f.write(fin_votar() + "\n")


if __name__ == "__main__":
    main(sys.argv[1], int(sys.argv[2]), int(sys.argv[3]), sys.argv[4])
