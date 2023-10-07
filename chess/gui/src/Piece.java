import java.util.ArrayList;

public abstract class Piece {
    Piece(Character N, Double V, Character C){
        this.Notation=N;
        this.Value=V;
        this.Color=C;
    }
    Character Notation;
    Double Value;
    Character Color;
    ArrayList<Integer> moves = new ArrayList<>(0);
    int numMoves;
    int boardInd;

    public abstract void moveGen();


}

class pawn extends Piece{
    pawn(Character N, Double V, Character C) {
        super(N,V,C);
    }
    public static final int WHITE_PAWN_DIRECTION = -1; // White pawns move up
    public static final int BLACK_PAWN_DIRECTION = 1;
    public void moveGen(){
        ArrayList<Integer> moves = new ArrayList<>();
        int dir = this.Color == 'W' ? WHITE_PAWN_DIRECTION : BLACK_PAWN_DIRECTION;
        if (this.numMoves > 0) {
            this.moves.add(this.boardInd + (dir * 8));
        } else {
            this.moves.add(this.boardInd + (dir * 8));
            this.moves.add(this.boardInd + (dir * 16));
        }
        for(int i : this.moves) {
            System.out.println(i);
        }
    }
}

class knight extends Piece{
    knight(Character N, Double V, Character C) {
        super(N,V,C);
    }

    public void moveGen() {
        System.out.println("nope");
    }
}
class bishop extends Piece{
    bishop(Character N, Double V, Character C) {
        super(N,V,C);
    }
    public void moveGen(){
        System.out.println("Nope");
    }

}
class rook extends Piece{
    rook(Character N, Double V, Character C) {
        super(N,V,C);
    }

    public void moveGen() {
        System.out.println("nope");
    }
}
class queen extends Piece{
    queen(Character N, Double V, Character C) {
        super(N,V,C);
    }

    public void moveGen() {
        System.out.println("nope");
    }
}
class king extends Piece {
    king(Character N, Double V, Character C) {
        super(N, V, C);
    }

    public void moveGen() {
        System.out.println("nope");
    }
}
