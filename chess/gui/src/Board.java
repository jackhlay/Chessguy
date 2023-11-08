import javax.swing.*;
import javax.swing.border.MatteBorder;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.util.ArrayList;

public class Board{
    private static int selected=-1;
    public static void go(){
        ArrayList<Piece> PieceList = new ArrayList<>();
        String pieces ="chess/images";
        char[][] BoardArr = Gamestate.Fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
        JPanel leftPanel = new JPanel();
        JFrame window = new JFrame();
        window.setLayout(new BorderLayout());
        JPanel Board = new JPanel(new GridLayout(8,8));

        for (int row = 0; row < 8; row++) {
            for (int col = 0; col < 8; col++) {
                JButton b = new JButton();
                b.setText(String.valueOf(BoardArr[row][col]));
                b.setForeground(new Color(85, 85, 85));
                b.setFocusPainted(false);
                b.setBorderPainted(false);
                //Colors: https://www.color-hex.com/color-palette/64811
                if ((row + col) % 2 == 1) b.setBackground(new Color(42, 34, 38));
                else {
                    b.setBackground(new Color(22, 24, 20));
                }
                b.setBorder(new MatteBorder(1, 1, 1, 1, new Color(85, 85, 85)));
                b.setVisible(true);
                b.setForeground(b.getBackground());
                int finalI = (8 * row + col);
                b.addActionListener(new ActionListener() {
                    @Override
                    public void actionPerformed(ActionEvent e) {
                        System.out.println("Selected: " + selected);
                        if(selected==-1){
                            System.out.println("Setting selected to " + finalI);
                            selected=finalI;
                        }
                        else if(selected==finalI){
                            System.out.println("Resetting from " + finalI + " to -1");
                            selected= -1;
                        }
                        else if(selected!=finalI){
                            System.out.println("Resetting from " + selected + " to -1");
                            selected=-1;
                        }
                    }
                });
                Board.add(b);
                String val = b.getText();
                switch (b.getText()){
                    case "P":
                        b.setIcon(new ImageIcon("chess/images/wp.png"));
                        Piece p = new pawn('P',1.0,'W');
                        p.boardInd = 8*row + col;
                        p.numMoves=0;
                        PieceList.add(p);
                        break;
                    case "K":
                        b.setIcon(new ImageIcon("chess/images/wK.png"));
                        Piece K = new king('K',0.0, 'W');
                        K.boardInd = 8*row + col;
                        K.numMoves = 0;
                        PieceList.add(K);
                        break;
                    case "N":
                        b.setIcon(new ImageIcon("chess/images/wN.png"));
                        break;
                    case "Q":
                        b.setIcon(new ImageIcon("chess/images/wQ.png"));
                        break;
                    case "B":
                        b.setIcon(new ImageIcon("chess/images/wB.png"));
                        break;
                    case "R":
                        b.setIcon(new ImageIcon("chess/images/wR.png"));
                        break;
                    case "p":
                        b.setIcon(new ImageIcon("chess/images/bp.png"));
                        break;
                    case "k":
                        b.setIcon(new ImageIcon("chess/images/bK.png"));
                        break;
                    case "n":
                        b.setIcon(new ImageIcon("chess/images/bN.png"));
                        break;
                    case "q":
                        b.setIcon(new ImageIcon("chess/images/bQ.png"));
                        break;
                    case "b":
                        b.setIcon(new ImageIcon("chess/images/bB.png"));
                        break;
                    case "r":
                        b.setIcon(new ImageIcon("chess/images/bR.png"));
                        break;
                    default:
                }

            }
        }
        leftPanel.setPreferredSize(new Dimension(300, 800)); // Adjust size as needed
        leftPanel.setBackground(Color.lightGray);

        Board.setPreferredSize(new Dimension(800,800));

        window.add(leftPanel,BorderLayout.WEST);
        window.setTitle("boby v0.23");
        window.add(Board,BorderLayout.CENTER);
        window.setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
        window.setSize(1100,800);
        window.setResizable(false);

        window.setVisible(true);
        Board.setVisible(true);
    }

    public static void main(String[] args) {
        go();
    }
}
