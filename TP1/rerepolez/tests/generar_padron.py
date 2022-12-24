import random

padron = [random.randint(1000000, 95000000) for i in range(300000)]
buscados = [random.choice(padron) for i in range(len(padron) // 5)]

with open("03_padron", "w") as f:
	for dni in padron:
		f.write(str(dni) + "\n")

with open("03_in", "w") as f:
	for dni in buscados:
		f.write("ingresar " + str(dni) + "\n")

with open("03_out", "w") as f:
	for i in range(len(buscados)):
		f.write("OK\n")
