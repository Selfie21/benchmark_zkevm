import csv
import time
import itertools
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from pprint import pprint

filename = "findings/data_zk.csv"
data = []
overhead = []
removed_rows = []

font = {'family' : 'normal',
        'size'   : 20}

matplotlib.rc('font', **font)

with open(filename, newline='') as csvfile:
    csvreader = csv.reader(csvfile)
    for i, row in enumerate(csvreader):
        if i % 5 == 0:
            removed_rows.append(row)
        elif i % 5 == 1:
            overhead.append(row)
        elif i % 5 == 4:
            data.append(row)

# Convert the data list into a numpy array
numpy_array_data = np.array(data).astype(float)
numpy_array_overhead = np.array(overhead).astype(float)
opcode_names = [removed_row[4][15:] for removed_row in removed_rows]


time_per_opcode = numpy_array_data[:, 3]
gas_used_total = numpy_array_data[:, 2] - numpy_array_overhead[:, 2]
gas_used_per_opcode = gas_used_total / numpy_array_data[:, 0]
time_per_gas = time_per_opcode / gas_used_per_opcode

array_rep_time = [None] * len(time_per_opcode)
array_rep_gas = [None] * len(time_per_opcode)
for i in range(len(time_per_opcode)):
    array_rep_time[i] = (opcode_names[i], time_per_opcode[i])
    array_rep_gas[i] = (opcode_names[i], time_per_gas[i])

array_rep_time = sorted(array_rep_time, key=lambda x: x[1], reverse=True)
array_rep_gas = sorted(array_rep_gas, key=lambda x: x[1], reverse=True)

colors = itertools.cycle(plt.cm.tab20.colors)
fig, ax = plt.subplots()
bars = ax.bar([x[0] for x in array_rep_time], [x[1] for x in array_rep_time])
for bar in bars:
    bar.set_color(next(colors))
ax.set_ylabel('Execution time [μs]')
ax.set_title('Per opcode execution time (overhead subtracted) native EVM')
plt.xticks(rotation=45, ha="right")  # Rotate x-axis labels for readability
plt.tight_layout()
plt.show()


colors = itertools.cycle(plt.cm.tab20.colors)
fig, ax = plt.subplots()
bars = ax.bar([x[0] for x in array_rep_gas], [x[1] for x in array_rep_gas])
for bar in bars:
    bar.set_color(next(colors))
ax.set_ylabel('Time per gas')
ax.set_title('Time per gas (overhead subtracted) native EVM')
plt.xticks(rotation=45, ha="right")  # Rotate x-axis labels for readability
plt.tight_layout()
plt.show()


prim = np.std([x[1] for x in array_rep_gas])
print(prim)