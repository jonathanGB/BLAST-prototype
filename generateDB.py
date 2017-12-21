# generate a database of 50 DNA sequences, each of length 50 to 100
# the nucleotides are chosen randomly

from random import randint, choice

db = open("db.txt", "w")

for line in range(0, 50):
  sequence = [choice(('A', 'C', 'T', 'G')) for _ in range(0, randint(50, 100))]
  sequence.append('\n')
  db.write("".join(sequence))

db.close()