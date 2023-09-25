import javax.swing.*;
import javax.swing.border.MatteBorder;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Board {

    public static void go(){
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
                if ((row+col) % 2 == 1) b.setBackground(new Color(42, 34, 38));
                else {
                    b.setBackground(new Color(22, 24, 20));
                }
                b.setBorder(new MatteBorder(1, 1, 1, 1, new Color(85, 85, 85)));
                b.setVisible(true);
                int finalI = (8*row + col);
                b.addActionListener(new ActionListener() {
                    @Override
                    public void actionPerformed(ActionEvent e) {
                        System.out.println(finalI);;
                    }
                });
                Board.add(b);
            }
        }
        leftPanel.setPreferredSize(new Dimension(300, 800)); // Adjust size as needed
        leftPanel.setBackground(Color.lightGray);

        Board.setPreferredSize(new Dimension(800,800));

        window.add(leftPanel,BorderLayout.WEST);
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
