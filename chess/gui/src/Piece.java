import org.w3c.dom.Notation;

import java.util.ArrayList;

public abstract class Piece {
    Piece(String T, Character N, Double V, Character C){
        this.PType=T;
        this.Notation=N;
        this.Value=V;
        this.Color=C;
    }
    String PType;
    Character Notation;
    Double Value;
    Character Color;

    public void Movegen() {}


}

class pawn extends Piece{
    pawn(String T, Character N, Double V, Character C) {
        super(T,N,V,C);
    }

}

class knight extends Piece{
    knight(String T, Character N, Double V, Character C) {
        super(T,N,V,C);
    }


}
class bishop extends Piece{
    bishop(String T, Character N, Double V, Character C) {
        super(T,N,V,C);
    }

}
class rook extends Piece{
    rook(String T, Character N, Double V, Character C) {
        super(T,N,V,C);
    }

}
class queen extends Piece{
    queen(String T, Character N, Double V, Character C) {
        super(T,N,V,C);
    }

}
class king extends Piece{
    king(String T, Character N, Double V, Character C) {
        super(T,N,V,C);
    }

}