import javax.swing.*;
import javax.swing.border.MatteBorder;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Board {

    public static void go(){
        JFrame window = new JFrame();
        JPanel Board = new JPanel(new GridLayout(8,8));

        for(int i=0; i<64; i++){
            JButton b = new JButton();
            if(i%2==1)b.setBackground(new Color(68,51,85));
            else{b.setBackground(new Color(153,153,153));}
            b.setBorder(new MatteBorder(1,1,1,1,new Color(85,85,85)));
            b.setVisible(true);
            int finalI = i;
            b.addActionListener(new ActionListener() {
                @Override
                public void actionPerformed(ActionEvent e) {
                    System.out.println(""+ finalI);
                }
            });
            Board.add(b);
        }
        window.add(Board);
        window.setSize(800,800);

        window.setVisible(true);
        Board.setVisible(true);
    }

    public static void main(String[] args) {
        go();
    }
}
