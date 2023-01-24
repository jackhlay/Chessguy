# ENGINE
# Handles gamestate info, and valid moves, writes gamelog.
import pygame

#images
#Black Pieces
bB = pygame.image.load("bB.png")
bK = pygame.image.load("bK.png")
bN = pygame.image.load("bN.png")
bp = pygame.image.load("bp.png")
bQ = pygame.image.load("bQ.png")
bR = pygame.image.load("bR.png")
#White Pieces
wB = pygame.image.load("wB.png")
wK = pygame.image.load("wK.png")
wN = pygame.image.load("wN.png")
wp = pygame.image.load("wp.png")
wQ = pygame.image.load("wQ.png")
wR = pygame.image.load("wR.png")

selected_square = None

wPiecesDict = {'KING': pygame.transform.scale(wK, (int(60*1.03), int(60*1.03))),
               'QUEEN': pygame.transform.scale(wQ, (int(60*1.03), int(60*1.03))),
               'BISHOP': pygame.transform.scale(wB, (int(60*1.03), int(60*1.03))),
               'KNIGHT': pygame.transform.scale(wN, (int(60*1.03), int(60*1.03))),
               'ROOK': pygame.transform.scale(wR, (int(60*1.03), int(60*1.03))),
               'PAWN': pygame.transform.scale(wp, (int(60*1.03), int(60*1.03)))}

bPiecesDict = {'KING': pygame.transform.scale(bK, (int(60*1.03), int(60*1.03))),
               'QUEEN': pygame.transform.scale(bQ, (int(60*1.03), int(60*1.03))),
               'BISHOP': pygame.transform.scale(bB, (int(60*1.03), int(60*1.03))),
               'KNIGHT': pygame.transform.scale(bN, (int(60*1.03), int(60*1.03))),
               'ROOK': pygame.transform.scale(bR, (int(60*1.03), int(60*1.03))),
               'PAWN': pygame.transform.scale(bp, (int(60*1.03), int(60*1.03)))}

class Space():
    occupied = False
    piece = None
    color = "White"

board = []
for i in range(64):
    board.append(Space())
    if i > 48:
        board[i].color="Black"

class GameState():
    def draw(self):
        pygame.init()
        board_size = (800, 800)
        screen = pygame.display.set_mode(board_size)
        white = (21, 17, 19)
        black = (11,12,10)
        pos = 0
        for i in range(8):
            for j in range(8):
                if (i+j) % 2 == 0:
                    color = black
                else:
                    color = white
                pygame.draw.rect(screen, color, (i*100, j*100, 100, 100))
                pos += 1
        #pygame.display.flip()
        running = True
        for it in range(64):

            if board[it].piece:
                if board[it].color=="Black":
                    screen.blit(bPiecesDict[board[it].piece], ((it%8)*100+20, (it//8)*100+20))
                    pygame.display.update()
                else:
                    screen.blit(wPiecesDict[board[it].piece], ((it%8)*100+20, (it//8)*100+20))
                    pygame.display.update()

        while running:
            for event in pygame.event.get():
                if event.type == pygame.MOUSEBUTTONDOWN:
                    x, y = event.pos
                    square_x, square_y = x // 100, y // 100
                    square_num = (square_y * 8) + square_x
                    takein(square_num)
                if event.type == pygame.QUIT:
                    running = False
        pygame.quit()


    def parfen(self, String):
        sqr = 0
        for i, char in enumerate(String):
            if char == "K":
                board[sqr].piece = "KING"
                board[sqr].color = "White"
                sqr+=1
            elif char == "k":
                board[sqr].piece = "KING"
                board[sqr].color = "Black"
                sqr+=1
            elif char == "Q":
                board[sqr].piece = "QUEEN"
                board[sqr].color = "White"
                sqr+=1
            elif char == "q":
                board[sqr].piece = "QUEEN"
                board[sqr].color = "Black"
                sqr+=1
            elif char == "B":
                board[sqr].piece = "BISHOP"
                board[sqr].color = "White"
                sqr+=1
            elif char == "b":
                board[sqr].piece = "BISHOP"
                board[sqr].color = "Black"
                sqr+=1
            elif char == "N":
                board[sqr].piece = "KNIGHT"
                board[sqr].color = "White"
                sqr+=1
            elif char == "n":
                board[sqr].piece = "KNIGHT"
                board[sqr].color = "Black"
                sqr+=1
            elif char == "R":
                board[sqr].piece = "ROOK"
                board[sqr].color = "White"
                sqr+=1
            elif char == "r":
                board[sqr].piece = "ROOK"
                board[sqr].color = "Black"
                sqr+=1
            elif char == "P":
                board[sqr].piece = "PAWN"
                board[sqr].color = "White"
                sqr+=1
            elif char == "p":
                board[sqr].piece = "PAWN"
                board[sqr].color = "Black"
                sqr+=1
            elif char == "/":
                continue
            else:
                sqr += int(char)
                i+=int(char)

def takein(ind):
    if board[ind].piece == "KING":
        print("single squares only")
    elif board[ind].piece == "QUEEN":
        print("UNTETHERED")
    elif board[ind].piece == "ROOK":
        print("slide moves NESW")
    elif board[ind].piece == "BISHOP":
        print("DIAGS ONLY")
    elif board[ind].piece == "KNIGHT":
        print("HORSEY")
    elif board[ind].piece == "PAWN":
        print("salt of the earth")

brd = GameState()
brd.parfen("1q6/1k1r2K1/1NP3pp/1P3N1Q/3Q2b1/nP6/4NR2/1P5B00")

brd.draw()