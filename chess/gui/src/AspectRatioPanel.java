import javax.swing.*;
import java.awt.*;

public class AspectRatioPanel extends JPanel {

    private double aspectRatio;

    public AspectRatioPanel(double aspectRatio) {
        this.aspectRatio = aspectRatio;
    }

    @Override
    public Dimension getPreferredSize() {
        Dimension parentSize = super.getPreferredSize();
        int width = parentSize.width;
        int height = parentSize.height;

        // Calculate the preferred size while maintaining the aspect ratio
        if (width > height) {
            width = (int) (height * aspectRatio);
        } else {
            height = (int) (width / aspectRatio);
        }

        return new Dimension(width, height);
    }

    public static void main(String[] args) {
        SwingUtilities.invokeLater(() -> {
            JFrame frame = new JFrame("Aspect Ratio JPanel");
            frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
            frame.setSize(400, 400);

            AspectRatioPanel panel = new AspectRatioPanel(16.0 / 9.0); // Set your desired aspect ratio here

            frame.add(panel);
            frame.setVisible(true);
        });
    }
}
