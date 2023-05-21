import pygame

# ENGINE
# Handles gamestate info, and valid moves, writes gamelog.
pygame.init()
pygame.display.set_caption('boby V0.228')
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

spots = [("a", 8), ("b", 8), ("c", 8), ("d", 8), ("e", 8), ("f", 8), ("g", 8), ("h", 8), ("a", 7), ("b", 7), ("c", 7),
         ("d", 7), ("e", 7), ("f", 7), ("g", 7), ("h", 7), ("a", 6), ("b", 6), ("c", 6), ("d", 6), ("e", 6), ("f", 6),
         ("g", 6), ("h", 6), ("a", 5), ("b", 5), ("c", 5), ("d", 5), ("e", 5), ("f", 5), ("g", 5), ("h", 5), ("a", 4),
         ("b", 4), ("c", 4), ("d", 4), ("e", 4), ("f", 4), ("g", 4), ("h", 4), ("a", 3), ("b", 3), ("c", 3), ("d", 3),
         ("e", 3), ("f", 3), ("g", 3), ("h", 3), ("a", 2), ("b", 2), ("c", 2), ("d", 2), ("e", 2), ("f", 2), ("g", 2),
         ("h", 2), ("a", 1), ("b", 1), ("c", 1), ("d", 1), ("e", 1), ("f", 1), ("g", 1), ("h", 1)]
         
board = []
piecearr = []

#Game Classes
class Space():
    occupied = False
    active = False
    place = None

class piece():
    type = None
    val = 0
    alive = True
    color = None
    moved = False
    nummoves = 0
    boardInd = None
    moves = []
    def movegen(self, ind):
        self.moves=[]
        if self.type == "ROOK":
            self.moves.extend(self.slides(ind))
        elif self.type == "KNIGHT":
            if board[ind].place[0] == "h":
                self.moves.extend([ind-17, ind-10, ind+6, ind+15])
            elif board[ind].place[0] == "a":
                self.moves.extend([ind-15, ind-6, ind+10, ind+17])
            else:
                self.moves.extend([ind-17, ind-15, ind-10, ind-6, ind+6, ind+10, ind+15, ind+17])
        elif self.type == "BISHOP":
            self.moves.extend(self.diags(ind))
        elif self.type == "QUEEN":
            self.moves.extend(self.slides(ind))
            self.moves.extend(self.diags(ind))
        elif self.type == "KING":
            self.moves = ["KCastle", "QCastle", ind-9, ind-8, ind-7, ind-1, ind+1, ind+7, ind+8, ind+9]
        elif self.type == "PAWN":
            if self.color == "White":
                takes = [ind-7, ind-9]
                if self.moved == False:
                    self.moves.extend([ind-8,ind-16])
                else:
                    self.moves.append(ind-8)

                for m in self.moves:
                    piece = next((piece for piece in piecearr if piece.boardInd == m), None)
                    if piece:
                        self.moves.remove(m)
                for t in takes:
                    piece = next((piece for piece in piecearr if piece.boardInd == t), None)
                    if piece and piece.color != self.color:
                        self.moves.append(t)

                if board[ind].place[1] == 5:
                    Lpassant = next((piece for piece in piecearr if piece.boardInd == ind-1), None)
                    Rpassant = next((piece for piece in piecearr if piece.boardInd == ind+1), None)
                    if board[ind-1].occupied and Lpassant.type == "PAWN" and Lpassant.nummoves == 1:
                        self.moves.append(ind-9)
                    if board[ind+1].occupied and Rpassant.type == "PAWN" and Rpassant.nummoves == 1:
                        self.moves.append(ind-7)
                    
            else:
                takes = [ind+7, ind+9]
                if self.moved == False:
                    self.moves.extend([ind+8,ind+16])
                else:
                    self.moves.append(ind+8)
                for m in self.moves:
                    piece = next((piece for piece in piecearr if piece.boardInd == m), None)
                    if piece:
                        self.moves.remove(m)
                for t in takes:
                    piece = next((piece for piece in piecearr if piece.boardInd == t), None)
                    if piece and piece.color != self.color:
                        self.moves.append(t)
                if board[ind].place[1] == 4:
                    Lpassant = next((piece for piece in piecearr if piece.boardInd == ind-1), None)
                    Rpassant = next((piece for piece in piecearr if piece.boardInd == ind+1), None)
                    if board[ind-1].occupied and Lpassant.type == "PAWN" and Lpassant.nummoves == 1:
                        self.moves.append(ind+7)
                    if board[ind+1].occupied and Rpassant.type == "PAWN" and Rpassant.nummoves == 1:
                        self.moves.append(ind+9)
    
        return self.moves
    
    def legals(self, arr):
        moves = []
        for m in arr:
            piece = next((piece for piece in piecearr if piece.boardInd == m), None)
            if m == "KCastle":
                if not self.moved:
                    if board[self.boardInd+1].occupied == False and board[self.boardInd+2].occupied == False:
                        moves.append("Kcastle")
            elif m == "QCastle":
                if not self.moved:
                    if board[self.boardInd-1].occupied == False and board[self.boardInd-2].occupied == False and board[self.boardInd-3].occupied == False:
                        moves.append("Qcastle")
            elif m >= 0 and m < 64:
                if not board[m].occupied:
                    moves.append(m)

                elif board[m].occupied and piece.color != self.color:
                    moves.append(m)
        return moves
    
    def slides(self, ind):
        moves = []
        for dir in range(4):
            for i in range (8):
                if dir == 0:#up
                    tp = ind - 8 * (i+1)
                elif dir == 1:#down
                    tp = ind + 8 * (i+1)
                elif dir == 2: #left
                    if tp % 8 == 0:
                        break
                    tp = ind - 1 * (i+1)
                elif dir == 3: #right
                    if tp % 8 == 7:
                        break
                    tp = ind + 1 * (i+1)
                p2 = next((piece for piece in piecearr if piece.boardInd == tp), None)
                if tp < 0 or tp > 63:
                    break
                if not p2:
                    moves.append(tp) 
                elif p2 and p2.color != self.color:
                    moves.append(tp)
                    break
                elif p2 and p2.color == self.color:
                    break
        return moves
    
    def diags(self, ind):
        moves = []
        for dir in range(4):
            for i in range(8):
                if dir == 0:  # up left
                    tp = ind - 9 * (i+1)
                elif dir == 1:  # up right
                    tp = ind - 7 * (i+1)
                elif dir == 2:  # down left
                    tp = ind + 7 * (i+1)
                    # print(tp, board[tp].occupied)
                elif dir == 3:  # down right
                    tp = ind + 9 * (i+1)
                    # print(tp, board[tp].occupied)

                p2 = next((piece for piece in piecearr if piece.boardInd == tp), None)
                if tp < 0 or tp > 63:
                    break

                if tp % 8 == 0 or tp % 8 == 7:
                    if not board[tp].occupied:
                        moves.append(tp)
                        break 
                    elif p2 and p2.color != self.color:
                        moves.append(tp)
                        break
                    elif p2 and p2.color == self.color:
                        break

                elif not board[tp].occupied:
                    moves.append(tp)
                elif p2 and p2.color != self.color:
                    moves.append(tp)
                    break
                elif p2 and p2.color == self.color:
                    break
        return moves

    def check(self, ind):
        OppMoves = []
        king = next((piece for piece in piecearr if piece.type == "KING" and piece.color == turn), None)
        for piece in piecearr:
            if piece.color != turn:
                OppMoves.extend(piece.movegen(piece.boardInd))
            if king.boardInd in OppMoves:
                return True
        return False
            
    def isLegalMove(self, ind):
        pass

    def makeMove(self, ind):
        pass

    def undoMove(self, ind):
        pass

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

        for m in piecearr:
            if m.type == "BISHOP":
                m.val = 3.1
            elif m.type == "KNIGHT":
                m.val == 3.0
            elif m.type == "ROOK":
                m.val = 5.0
            elif m.type == "QUEEN":
                m.val = 9.0
            elif m.type == "PAWN":
                m.val = 1.0
            elif m.type == "KING":
                m.val = 0

#Draws board
def draw():
    global turn,screen
    pygame.init()
    board_size = (800, 800)
    screen = pygame.display.set_mode(board_size)
    light = (42, 34, 38)
    dark = (22,24,20)
    pos = 0
    for i in range(8): #draw board
        for j in range(8):
            if (i+j) % 2 == 0:
                color = dark
            else:
                color = light
            pygame.draw.rect(screen, color, (i*100, j*100, 100, 100))
            pos += 1

    for i in range(64): #draw pieces
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
                moves = []
                x, y = event.pos
                ind = ((y//100) * 8) + (x//100)
                orig = takein(x,y)
                piece = next((piece for piece in piecearr if piece.boardInd == ind), None)
                print(ind)
                if not orig.occupied or piece.color != turn: #If you click on an empty square or the wrong color
                    drawit(screen)
                    continue

                elif piece.check(ind): #Check evaluation
                    print("IN CHECK")
                    allmoves = []
                    legalmoves = []
                    for p in piecearr:
                        if p.color == turn:
                            allmoves += piece.movegen(piece.boardInd)
                    for move in allmoves:
                        if piece.isLegalMove(move):
                            legalmoves.append(move)

                else: #If piece is not in check
                    dragging=True
                    moves = piece.legals(piece.movegen(ind))
                    print(moves)
                    if piece.color=="White":
                        img = wPiecesDict[piece.type]
                    else:
                        img = bPiecesDict[piece.type]
                

            if event.type== pygame.MOUSEBUTTONUP: 
                dragging = False
                x, y = event.pos
                ind2 = ((y//100) * 8) + (x//100)
                piece2 = next((piece for piece in piecearr if piece.boardInd == ind2), None)
                fin = takein(x,y)
                if not piece or fin == orig or turn != piece.color: #If no piece or same piece or wrong color
                    drawit(screen)
                
                elif board[ind2].occupied and piece2.color == piece.color: #If piece of same color
                    drawit(screen)
                
                elif piece.type == "KING" and 1 < abs(ind2-ind) < 5: #Castline Block
                    if "KCastle" or "QCastle" in piece.moves:
                        if ind2-ind > 0:
                            rook = next((piece for piece in piecearr if piece.boardInd == ind+3), None)
                            piece.boardInd += 2
                            orig.occupied = False
                            orig.active = False
                            board[rook.boardInd].occupied = False
                            board[rook.boardInd].active = False
                            rook.boardInd -= 2
                            board[piece.boardInd].occupied = True
                            board[piece.boardInd].active = True
                            board[rook.boardInd].occupied = True
                            board[rook.boardInd].active = True
                            piece.moved = True
                            rook.moved = True
                            piece.nummoves += 1
                            rook.nummoves += 1
                            
                        elif ind2-ind < 0:
                            rook = next((piece for piece in piecearr if piece.boardInd == ind-4), None)
                            piece.boardInd -= 2
                            orig.occupied = False
                            orig.active = False
                            board[rook.boardInd].occupied = False
                            board[rook.boardInd].active = False
                            rook.boardInd += 3
                            board[piece.boardInd].occupied = True
                            board[piece.boardInd].active = True
                            board[rook.boardInd].occupied = True
                            board[rook.boardInd].active = True
                            piece.moved = True
                            rook.moved = True
                            piece.nummoves += 1
                            rook.nummoves += 1
                    
                    turn = "White" if turn == "Black" else "Black"
                    drawit(screen)

                elif ind2 in moves: #Check if move is legal
                    if piece.type == "PAWN" and (abs(ind2-ind) == 9):
                        if turn == "White" and board[ind2+8].occupied and board[ind].place[1]==5: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2+8), None)
                            if piece2.color == "Black":
                                orig.occupied = False
                                orig.active = False
                                piece.boardInd = ind2
                                piece.moved = True
                                piece.nummoves += 1
                                fin.occupied = True
                                fin.active = True
                                board[ind2+8].occupied = False
                        elif turn == "Black" and board[ind2-8].occupied and board[ind].place[1]==4: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2-8), None)
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True
                            board[ind2-8].occupied = False
                            
                        else:
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True

                    elif piece.type == "PAWN" and (abs(ind2-ind) == 7): #Grand En Passant block
                        if turn == "White" and board[ind2+8].occupied and board[ind].place[1]==5: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2+8), None)
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True
                            board[ind2+8].occupied = False
                        elif turn == "Black" and board[ind2-8].occupied and board[ind].place[1]==4: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2-8), None)
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True
                            board[ind2-8].occupied = False
                        else:
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True

                    elif turn == "White" and piece.type == "PAWN" and board[ind2].place[1] == 8: #Promotion block
                        orig.occupied = False
                        orig.active = False
                        piece.type = "QUEEN"
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True

                    elif turn == "Black" and piece.type == "PAWN" and board[ind2].place[1] == 1: #Promotion block
                        orig.occupied = False
                        orig.active = False
                        piece.type = "QUEEN"
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True    


                    else:
                        orig.occupied = False
                        orig.active = False
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True

                    if piece2:
                        print(f"Piece Array Length {len(piecearr)}")
                        piece2.alive = False 
                        piece2.boardInd = -1
                        piecearr.remove(piece2)
                        print(f"Piece Array Length {len(piecearr)}")


                    turn = "White" if turn == "Black" else "Black"
                    drawit(screen)      

            if event.type == pygame.MOUSEMOTION:
                if dragging:
                    clock.tick(60)
                    x, y = event.pos
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
    
    [#previous print tests
    # print('Active: {}'.format(spot.active))
    # print(spot.color)
    # if spot.occupied:
    #     if piece.type == "KING":
    #         print(f"{piece.color} KING")
    #     elif piece.type == "QUEEN":
    #         print(f"{piece.color} QUEEN")
    #     elif piece.type == "ROOK":
    #         print(f"{piece.color} ROOK")
    #     elif piece.type == "BISHOP":
    #         print(f"{piece.color} BISHOP")
    #     elif piece.type == "KNIGHT":
    #         print(f"{piece.color}HORSEY")
    #     elif piece.type == "PAWN":
    #         print(f"{piece.color} PAWN")
    # else:
    #     print("empty")
    ]

    return spot

#GO! GO! GO!             
def start(string="rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"):
    for i in range(64):
        board.append(Space())
        board[i].place = spots[i]
    
    parfen(string)
    draw()
    go(screen)

start()
#default string: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
#diagonal moves test string: rnbqkbnr/p6p/8/8/8/8/P6P/RNBQKBNR
#Castle test string: r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R