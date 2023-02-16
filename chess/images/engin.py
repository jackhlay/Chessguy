import sys
import pygame

# ENGINE
# Handles gamestate info, and valid moves, writes gamelog.
pygame.init()
pygame.display.set_caption('boby V0.22')
turn ="White"

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


#Game Classes
class Space():
    occupied = False
    active = False
    piece = None
    color = "White"
    place = None


#Gamestate Variables block

spots = [("a", 8), ("b", 8), ("c", 8), ("d", 8), ("e", 8), ("f", 8), ("g", 8), ("h", 8), ("a", 7), ("b", 7), ("c", 7),
         ("d", 7), ("e", 7), ("f", 7), ("g", 7), ("h", 7), ("a", 6), ("b", 6), ("c", 6), ("d", 6), ("e", 6), ("f", 6),
         ("g", 6), ("h", 6), ("a", 5), ("b", 5), ("c", 5), ("d", 5), ("e", 5), ("f", 5), ("g", 5), ("h", 5), ("a", 4),
         ("b", 4), ("c", 4), ("d", 4), ("e", 4), ("f", 4), ("g", 4), ("h", 4), ("a", 3), ("b", 3), ("c", 3), ("d", 3),
         ("e", 3), ("f", 3), ("g", 3), ("h", 3), ("a", 2), ("b", 2), ("c", 2), ("d", 2), ("e", 2), ("f", 2), ("g", 2),
         ("h", 2), ("a", 1), ("b", 1), ("c", 1), ("d", 1), ("e", 1), ("f", 1), ("g", 1), ("h", 1)]
board = []
for i in range(64):
    board.append(Space())
    board[i].place = spots[i]
    if i > 48:
        board[i].color="Black"


#Functions Block
def parfen(String):
    sqr = 0
    for i in range(64):
        board[i].color = None

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
    draw()

def draw():
    global turn
    pygame.init()
    board_size = (800, 800)
    screen = pygame.display.set_mode(board_size)
    light = (42, 34, 38)
    dark = (22,24,20)
    pos = 0
    for i in range(8):
        for j in range(8):
            if (i+j) % 2 == 0:
                color = dark
            else:
                color = light
            pygame.draw.rect(screen, color, (i*100, j*100, 100, 100))
            pos += 1

    dragging_piece=None
    running = True
    for i in range(64):
        if board[i].piece:
            board[i].active=True
            if board[i].color=="Black":
                screen.blit(bPiecesDict[board[i].piece], ((i%8)*100+20, (i//8)*100+20))
                pygame.display.flip()
            else:
                screen.blit(wPiecesDict[board[i].piece], ((i%8)*100+20, (i//8)*100+20))
                pygame.display.flip()


    dragging=False
    offset_x, offset_y = 0, 0
    clock = pygame.time.Clock()

    while running:

        pygame.init()
        for event in pygame.event.get():
            if event.type == pygame.MOUSEBUTTONDOWN:
                x, y = event.pos
                orig = takein(x,y)
                if orig.active:
                    dragging=True
                    if orig.color=="White":
                        img = wPiecesDict[orig.piece]
                    else:
                        img = bPiecesDict[orig.piece]

            if event.type== pygame.MOUSEBUTTONUP:
                dragging = False
                x, y = event.pos
                fin = takein(x,y)
                if fin == orig or turn != orig.color:
                    continue
                else:
                    fin.piece = orig.piece
                    fin.color=orig.color
                    fin.active=True
                    orig.piece=None
                    orig.color=None
                    orig.active=False
                    if turn == "White":
                        turn = "Black"
                    else:
                        turn="White"
                    draw()

            if event.type == pygame.MOUSEMOTION:
                if dragging:
                    clock.tick(1000)
                    x, y = event.pos
                    imgx = x + offset_x
                    imgy = y + offset_y
                    screen.blit(img, (x,y))
                    pygame.display.update()

            if event.type == pygame.QUIT:
                running = False
                sys.exit()

def takein(x,y):
    square_x, square_y = x // 100, y // 100
    ind = (square_y * 8) + square_x
    spot = board[ind]

    # print('Active: {}'.format(spot.active))
    # print(spot.color)

    return spot

    # if spot.piece == "KING":
    #     print("you, sir")
    # elif spot.piece == "QUEEN":
    #     print("my freakin wife")
    # elif spot.piece == "ROOK":
    #     print("slide moves NESW")
    # elif spot.piece == "BISHOP":
    #     print("diags only")
    # elif spot.piece == "KNIGHT":
    #     print("HORSEY")
    # elif spot.piece == "PAWN":
    #     print("7 Spares")
    # else:
    #     print("empty")


# def movegen(ind):
#     spot = board[ind]
#     if spot.piece == "KING":
#         pos=[ind-9, ind-8, ind-7, ind-1, ind+1, ind+7, ind+8, ind+9]
#         #moves=isLegal(pos)
#     elif spot.piece == "QUEEN":
#         pos = []
#         #moves=isLegal(pos)
#     elif spot.piece == "ROOK":
#         pos=rook(ind)
#         #moves=isLegal(pos)
#     elif spot.piece == "BISHOP":
#         pos = []
#         #moves=isLegal(pos)
#     elif spot.piece == "KNIGHT":
#         pos = [ind-17, ind-15, ind-10, ind-6, ind+6, ind+10, ind+15, ind+17]
#         #moves=isLegal(pos)
#     elif spot.piece == "PAWN":
#         if spot.color == "White":
#             pos = [ind-8]
#             #moves=isLegal(pos)
#         elif spot.color == "Black":
#             pos = [ind+8]
#             #moves=isLegal(pos)
#     #return moves
#
# def rook(ind):
#     moves=[]
#     for i in range(-8,8):
#         if board[i].occupied:
#             if board[i].color==board[ind].color:
#                 break
#             else:
#                 moves.append(i)
#         else:
#             moves.append(i)
#     for i in range(moves.size()):
#         moves[i]=board[moves[i].place]
#     return moves

def start(string):
    for i in board:
        i.piece=None
    parfen(string)

start("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")