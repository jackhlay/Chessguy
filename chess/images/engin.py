# ENGINE
# Handles gamestate info, and valid moves, writes gamelog.
import pygame

class Space():
    def __init__(self, clr, piece, occupied):
        self.color = clr
        self.status = occupied
        if occupied:
            self.piece = piece
        else:
            self.piece = None

class GameState():
    def draw(self):
        pygame.init()
        board = [Space]*64
        board_size = (800, 800)
        screen = pygame.display.set_mode(board_size)
        for i in range(64):
            board[i] = i
        white = (15, 17, 17)
        black = (0,0,0)
        pos = 0
        for i in range(8):
            for j in range(8):
                if (i+j) % 2 == 0:
                    color = black
                else:
                    color = white
                pygame.draw.rect(screen, color, (i*100, j*100, 100, 100))
                pos += 1
        pygame.display.flip()
        running = True
        while running:
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    running = False
        pygame.quit()
        print(board)

    def parfen(self, String):
        sqr = 0;
        for i, char in enumerate(String):
            if char == "K" or char == "k":
                sqr+=1
                print("KING\n")
            elif char == "Q" or char == "q":
                sqr+=1
                print("QUEEN\n")
            elif char == "B" or char == "b":
                sqr+=1
                print("BISHOP\n")
            elif char == "N" or char == "n":
                sqr+=1
                print("KNIGHT\n")
            elif char == "R" or char == "r":
                sqr+=1
                print("ROOK\n")
            elif char == "P" or char == "p":
                sqr+=1
                print("PAWN\n")
            elif char == "/":
                continue
            else:
                sqr+=ord(char)
                print(sqr)

board = GameState()
board.draw()