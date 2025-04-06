const N = 8;
let board : number[][] = [];

function initBoard(){
    for (let i = 0; i < N; i++) {
        board[i] = [];
        for (let j = 0; j < N; j++) {
            board[i][j] = 0;
        }
    }
}




function isSafe(board : number[][], row : number, col : number){
    for (let i = 0; i < col; i++) {
        if(board[row][i] == 1)return false;
    }
    for (let i = row, j = col; i >= 0 && j >= 0; i--, j--) {
        if(board[i][j] == 1)return false;
    }
    for (let i = row, j = col; i < N && j >= 0; i++, j--) {
        if(board[i][j] == 1)return false;
    }
    return true;
}


function solveNQueens(){
    initBoard();
    if(!solveNQueensUtil(board, 0))return false;
    printBoard(board);
}


function solveNQueensUtil(board : number[][], col : number){
    if(col == N)return true;
    for (let i = 0; i < N; i++) {
        if(isSafe(board, i, col)){
            board[i][col] = 1;
            if(solveNQueensUtil(board, col + 1))return true;
            board[i][col] = 0;
        }
    }
    return false;
}




function printBoard(board : number[][]){
    for (let i = 0; i < N; i++) {
        for (let j = 0; j < N; j++) {
            console.log(board[i][j] == 1? "Q ": ". ");
        }
        console.log();
    }
}




solveNQueens();