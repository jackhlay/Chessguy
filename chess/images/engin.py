import pygame

# ENGINE
# Handles gamestate info, and valid moves, writes gamelog.
pygame.init()
pygame.display.set_caption('boby V0.221')
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

piecearr = []

#Game Classes
class Space():
    occupied = False
    active = False
    place = None

class piece():
    type = None
    alive = True
    color = None
    moved = False
    boardInd = None
    moves = []
    def movegen(self, ind):
        self.moves=[]
        if self.type == "ROOK":
            pass
        elif self.type == "KNIGHT":
            pass
        elif self.type == "BISHOP":
            pass
        elif self.type == "QUEEN":
            pass
        elif self.type == "KING":
            pass
        elif self.type == "PAWN":
            if self.color == "White":
                if self.moved == False:
                    self.moves.extend([ind-8,ind-16])
                else:
                    self.moves.append(ind-8)
            else:
                if self.moved == False:
                    self.moves.extend([ind+8,ind+16])
                else:
                    self.moves.append(ind+8)
        return self.moves

#Functions Block

#Parser
def parfen(String):
    sqr = 0
    for i in range(64):
        board[i].occupied = False

    for i, char in enumerate(String):
        if char == "K":
            WKing = piece()
            WKing.color = "White"
            WKing.type = "KING"
            WKing.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WKing)
            sqr+=1
        elif char == "k":
            BKing = piece()
            BKing.color = "Black"
            BKing.type = "KING"
            BKing.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BKing)
            sqr+=1
        elif char == "Q":
            WQueen = piece()
            WQueen.color = "White"
            WQueen.type = "QUEEN"
            WQueen.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WQueen)
            sqr+=1
        elif char == "q":
            BQueen = piece()
            BQueen.color = "Black"
            BQueen.type = "QUEEN"
            BQueen.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BQueen)
            sqr+=1
        elif char == "B":
            WBishop = piece()
            WBishop.color = "White"
            WBishop.type = "BISHOP"
            WBishop.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WBishop)
            sqr+=1
        elif char == "b":
            BBishop = piece()
            BBishop.color = "Black"
            BBishop.type = "BISHOP"
            BBishop.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BBishop)
            sqr+=1
        elif char == "N":
            WKnight = piece()
            WKnight.color = "White"
            WKnight.type = "KNIGHT"
            WKnight.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WKnight)
            sqr+=1
        elif char == "n":
            BKnight = piece()
            BKnight.color = "Black"
            BKnight.type = "KNIGHT"
            BKnight.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BKnight)
            sqr+=1
        elif char == "R":
            WRook = piece()
            WRook.color = "White"
            WRook.type = "ROOK"
            WRook.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WRook)
            sqr+=1
        elif char == "r":
            BRook = piece()
            BRook.color = "Black"
            BRook.type = "ROOK"
            BRook.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BRook)
            sqr+=1
        elif char == "P":
            WPawn = piece()
            WPawn.color = "White"
            WPawn.type = "PAWN"
            WPawn.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WPawn)
            sqr+=1
        elif char == "p":
            BPawn = piece()
            BPawn.color = "Black"
            BPawn.type = "PAWN"
            BPawn.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BPawn)
            sqr+=1
        elif char == "/":
            continue
        else:
            sqr += int(char)
            i+=int(char)

#Draws board
def draw():
    global turn,screen
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

    for i in range(64):
        if board[i].occupied:
            piece = next((piece for piece in piecearr if piece.boardInd == i), None)
            if piece.color=="Black":
                screen.blit(bPiecesDict[piece.type], ((i%8)*100+20, (i//8)*100+20))
                pygame.display.flip()
            else:
                screen.blit(wPiecesDict[piece.type], ((i%8)*100+20, (i//8)*100+20))
                pygame.display.flip()

#Game loop and maintainence functions
def go(screen):
    global turn
    running = True
    dragging=False
    offset_x, offset_y = 0, 0
    clock = pygame.time.Clock()
    pygame.init()
    #clean up this while loop
    while running:
        
        for event in pygame.event.get():
            if event.type == pygame.MOUSEBUTTONDOWN:
                x, y = event.pos
                square_x, square_y = x // 100, y // 100
                ind = (square_y * 8) + square_x
                orig = takein(x,y)
                piece = next((piece for piece in piecearr if piece.boardInd == ind), None)
                piece.movegen(ind)
                if orig.occupied and piece.color == turn:
                    dragging=True
                    if piece.color=="White":
                        img = wPiecesDict[piece.type]
                    else:
                        img = bPiecesDict[piece.type]
                print(piece.moves)
                

            if event.type== pygame.MOUSEBUTTONUP:
                dragging = False
                x, y = event.pos
                ind = (square_y * 8) + square_x
                # fin = takein(x,y)
                # if fin == orig or turn != piece.color:
                #     drawit(screen)
                
                drawit(screen)


            if event.type == pygame.MOUSEMOTION:
                if dragging:
                    clock.tick(60)
                    x, y = event.pos
                    imgx = x + offset_x
                    imgy = y + offset_y
                    drawit(screen)
                    screen.blit(img, (x,y))
                    pygame.display.update()

            if event.type == pygame.QUIT:
                running = False
                exit()

def drawit(screen):
    light = (42, 34, 38)
    dark = (22,24,20)
    pos = 0

    foreground = pygame.Surface((800,800))
    background = pygame.Surface((800,800))
    background.blit(screen, (0, 0))

    for i in range(8):
        for j in range(8):
            if (i+j) % 2 == 0:
                color = dark
            else:
                color = light
            pygame.draw.rect(foreground, color, (i*100, j*100, 100, 100))
            pos += 1

    for i in range(64):
        if board[i].occupied:
            piece = next((piece for piece in piecearr if piece.boardInd == i), None)
            if piece.color=="Black":
                foreground.blit(bPiecesDict[piece.type], ((i%8)*100+20, (i//8)*100+20))
            else:
                foreground.blit(wPiecesDict[piece.type], ((i%8)*100+20, (i//8)*100+20))
    
    background.blit(foreground, (0, 0))
    screen.blit(background, (0, 0))
    pygame.display.flip()

def takein(x,y):
    square_x, square_y = x // 100, y // 100
    ind = (square_y * 8) + square_x
    spot = board[ind]
    piece = next((piece for piece in piecearr if piece.boardInd == ind), None)
    
    #previous print tests
    # print('Active: {}'.format(spot.active))
    # print(spot.color)
    if spot.occupied:
        if piece.type == "KING":
            print(f"{piece.color} KING")
        elif piece.type == "QUEEN":
            print(f"{piece.color} QUEEN")
        elif piece.type == "ROOK":
            print(f"{piece.color} ROOK")
        elif piece.type == "BISHOP":
            print(f"{piece.color} BISHOP")
        elif piece.type == "KNIGHT":
            print("HORSEY")
        elif piece.type == "PAWN":
            print(f"{piece.color} PAWN")
    else:
        print("empty")

    return spot

#GO! GO! GO!                  
def start(string="rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"):
    global board
    board,spots = [],[("a", 8), ("b", 8), ("c", 8), ("d", 8), ("e", 8), ("f", 8), ("g", 8), ("h", 8), ("a", 7), ("b", 7), ("c", 7),
         ("d", 7), ("e", 7), ("f", 7), ("g", 7), ("h", 7), ("a", 6), ("b", 6), ("c", 6), ("d", 6), ("e", 6), ("f", 6),
         ("g", 6), ("h", 6), ("a", 5), ("b", 5), ("c", 5), ("d", 5), ("e", 5), ("f", 5), ("g", 5), ("h", 5), ("a", 4),
         ("b", 4), ("c", 4), ("d", 4), ("e", 4), ("f", 4), ("g", 4), ("h", 4), ("a", 3), ("b", 3), ("c", 3), ("d", 3),
         ("e", 3), ("f", 3), ("g", 3), ("h", 3), ("a", 2), ("b", 2), ("c", 2), ("d", 2), ("e", 2), ("f", 2), ("g", 2),
         ("h", 2), ("a", 1), ("b", 1), ("c", 1), ("d", 1), ("e", 1), ("f", 1), ("g", 1), ("h", 1)]
    for i in range(64):
        board.append(Space())
        board[i].place = spots[i]
    
    parfen(string)
    draw()
    go(screen)

start()