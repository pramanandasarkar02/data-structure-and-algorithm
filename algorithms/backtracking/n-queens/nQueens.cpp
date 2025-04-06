#include <iostream>
#include <vector>
#include <string>
using namespace std;


int N = 8;

vector<vector<int>> board = vector<vector<int>>(N, vector<int>(N, 0));


bool isSafe(int row, int col) {
    for (int i = 0; i < col; i++) {
        if (board[row][i] == 1) return false;
    }
    for (int i = row, j = col; i >= 0 && j >= 0; i--, j--) {
        if (board[i][j] == 1) return false;
    }
    for (int i = row, j = col; i < N && j >= 0; i++, j--) {
        if (board[i][j] == 1) return false;
    }
    return true;
}


bool solveNQueensUtil(int col) {
    if (col == N) {
        return true;
    }
    for (int i = 0; i < N; i++) {
        if (isSafe(i, col)) {
            board[i][col] = 1;
            if (solveNQueensUtil(col + 1)) return true;
            board[i][col] = 0;
        }
    }
    return false;
}

bool solveNQueens() {
    return solveNQueensUtil(0);
}


void printBoard() {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N; j++) {
            if (board[i][j] == 1) cout << "Q ";
            else cout << ". ";
        }
        cout << endl;
    }
}


int main() {
    solveNQueens();
    printBoard();
}