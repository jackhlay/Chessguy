import java.util.ArrayList;

public class Gamestate {
    char[][] Board = new char[8][8];

    public static char[][] Fen(String x) {
        char[][] Board = new char[8][8];
        char[] word = x.toCharArray();
        int r = 0;
        int c = 0;

        for (int i = 0; i < word.length; i++) {
            char currentChar = word[i];
            if (Character.isDigit(currentChar)) {
                // If the character is a digit, skip that number of columns.
                int numEmpty = Character.getNumericValue(currentChar);
                for (int j = 0; j < numEmpty; j++) {
                    Board[r][c] = ' '; // Assuming space represents an empty square
                    c++;
                }
            } else if (currentChar == '/') {
                // If the character is "/", move to the next row.
                r++;
                c = 0;
            } else {
                // Otherwise, it's a chess piece notation.
                Board[r][c] = currentChar;
                c++;
            }
            if (r >= 8) {
                break;
            }
        }
        return Board;
    }

    public static void main(String[] args) {
        char[][] Board = Fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
        for (int row = 0; row < Board.length; row++) {
            for (int col = 0; col < Board[row].length; col++) {
                System.out.print(Board[row][col] + " ");
            }
            System.out.println(); // Move to the next line after printing each row
        }
    }
}
