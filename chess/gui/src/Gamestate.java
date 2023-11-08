import java.util.ArrayList;

public class Gamestate {
    char turn = 'W';
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

}