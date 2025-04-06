N = 8

board = [[0 for _ in range(N)] for _ in range(N)]



def NQueens():
    if solveNQueens():
        printBoard(board)
    else:
        print("No Solution")

def solveNQueens():
    return solveNQueensUtil(0)

def solveNQueensUtil(col):
    if col == N:
        return True

    for i in range(N):
        if isSafe(i, col):
            board[i][col] = 1
            if solveNQueensUtil(col + 1):
                return True
            board[i][col] = 0
    return False

def isSafe(row, col):
    for i in range(col):
        if board[row][i] == 1:
            return False
    for i, j in zip(range(row, -1, -1), range(col, -1, -1)):
        if board[i][j] == 1:
            return False
    for i, j in zip(range(row, N, 1), range(col, -1, -1)):
        if board[i][j] == 1:
            return False
    return True

def printBoard(board):
    for row in  board:
        for val in row:
            print("Q" if val == 1 else ".", end=" ")
        print()




NQueens()