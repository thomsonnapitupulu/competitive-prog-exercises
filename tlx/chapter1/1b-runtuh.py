R, C = [int (a) for a in input().split()]

fulfilled_rows = []
index_rows = [list() for f in range(C)]
result_matrix = [[0]*C for g in range(R)]

for x in range(R):
    row = input()
    
    col_counter = 0
    fulfilled_counter = 0

    for item in row:
        if int(item) == 1:
            index_rows[col_counter].append(x)
            fulfilled_counter += 1
            result_matrix[x][col_counter] = 1
        col_counter += 1

        if fulfilled_counter == C:
            fulfilled_rows.append(x)

while len(fulfilled_rows) > 0:
    result_matrix = [[0]*C for g in range(R)]
    r_count = 0

    for row_indexes in index_rows:
        substractor = 0
        total_removed_rows = len(fulfilled_rows)

        for i in range(len(row_indexes)-1, -1, -1):
            if fulfilled_rows[len(fulfilled_rows) -1] == row_indexes[i]:
                break

            result_matrix[row_indexes[i]][r_count] = 1
            substractor += 1

        diff = 0
        if i < (len(row_indexes)-1): # cek apakah sudah melakukan substraction
            diff = (row_indexes[i+1] - fulfilled_rows[len(fulfilled_rows) -1]) - 1

        if row_indexes[len(row_indexes)-1] == fulfilled_rows[len(fulfilled_rows) -1]:
            start_index = R-1
        else:
            start_index = fulfilled_rows[len(fulfilled_rows) -1] + diff
    
        remaining_row_indexes_counter = len(row_indexes) - substractor - len(fulfilled_rows)
        last_index = start_index - remaining_row_indexes_counter

        for i in range(start_index, last_index, -1):
            result_matrix[i][r_count] = 1

        r_count += 1

    index_rows.clear()
    index_rows = [list() for f in range(C)]
    fulfilled_rows.clear()
    row_index = 0
    for result_item in result_matrix:
        col_counter = 0
        fulfilled_counter = 0
        for cell in result_item:
            if cell == 1:
                index_rows[col_counter].append(row_index)
                fulfilled_counter += 1
            col_counter += 1

            if fulfilled_counter == C:
                fulfilled_rows.append(row_index)

        row_index += 1

    if len(fulfilled_rows) > 0:
        result_matrix.clear()

for result_item in result_matrix:
    print("".join(str(n) for n in result_item))
