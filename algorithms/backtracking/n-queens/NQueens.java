public class NQueens{
    static int N = 8;
    static int[][] board ;

    public static void main(String[] args) {
        for (int i = 0; i < 20; i++){
           
            N = i;
            board = new int[N][N];

            if(solveNQueens()){
                System.out.println("Solution for N = "+ i);
                printSolution();
            }
            else{
                System.out.println("No Solution for N = "+ i);
            }
        }
        // solveNQueens();
        // printSolution();

        
    }


    public static boolean solveNQueens(){
        return solve(0);
    }

    private static boolean solve(int col){
        if(col == N){
            return true;
        }
        for (int row = 0; row < N; row++) {
            if(isSafe(row, col)){
                board[row][col] = 1;
                if(solve( col + 1)) return true;
                board[row][col] = 0;
            }
        }
        return false;
    }

    private static  boolean isSafe(int row, int col){
        for(int j = 0; j < col; j++){
            if(board[row][j] == 1)return false;
        }
        for(int i = row, j = col; i >= 0 && j >= 0; i--, j--){
            if(board[i][j] == 1)return false;
        }
        for(int i = row, j = col; i < N && j >= 0; i++, j--){
            if(board[i][j] == 1)return false;
        }


        return true;
    }

    public static void printSolution(){
        for(int i = 0; i < N; i++){
            for(int j = 0; j < N; j++){
                System.out.print(board[i][j] == 1? "Q ": ". ");
            }
            System.out.println();
        }
    }

}