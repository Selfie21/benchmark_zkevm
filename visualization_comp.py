import csv
import time
import itertools
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from pprint import pprint

font = {'family' : 'normal',
        'size'   : 20}
matplotlib.rc('font', **font)

def NormalizeData(data):
    return (data - np.min(data)) / (np.max(data) - np.min(data))


filename = "findings/data_nat.csv"
data = []
overhead = []
removed_rows = []

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

time_per_opcode = NormalizeData(time_per_opcode)
time_per_gas = NormalizeData(time_per_gas)

array_rep_time = [None] * len(time_per_opcode)
array_rep_gas = [None] * len(time_per_opcode)
for i in range(len(time_per_opcode)):
    array_rep_time[i] = (opcode_names[i], time_per_opcode[i])
    array_rep_gas[i] = (opcode_names[i], time_per_gas[i])

array_rep_time = sorted(array_rep_time, key=lambda x: x[1], reverse=True)
array_rep_gas = sorted(array_rep_gas, key=lambda x: x[1], reverse=True)


fig, ax = plt.subplots()
ax.bar([x[0] for x in array_rep_gas], [x[1] for x in array_rep_gas], label='native EVM')




filename = "findings/data_zk.csv"
data = []
overhead = []
removed_rows = []

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

time_per_opcode = NormalizeData(time_per_opcode)
time_per_gas = NormalizeData(time_per_gas)

array_rep_time = [None] * len(time_per_opcode)
array_rep_gas = [None] * len(time_per_opcode)
for i in range(len(time_per_opcode)):
    array_rep_time[i] = (opcode_names[i], time_per_opcode[i])
    array_rep_gas[i] = (opcode_names[i], time_per_gas[i])

array_rep_time = sorted(array_rep_time, key=lambda x: x[1], reverse=True)
array_rep_gas = sorted(array_rep_gas, key=lambda x: x[1], reverse=True)




ax.bar([x[0] for x in array_rep_gas], [x[1] for x in array_rep_gas], width=1/2, label='zkEVM')
ax.set_ylabel('Time per gas')
ax.set_title('Time per gas (overhead subtracted) normalized')
ax.legend()
plt.xticks(rotation=45, ha="right")
plt.tight_layout()
plt.show()
