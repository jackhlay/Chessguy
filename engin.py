import pygame

pygame.init()
pygame.display.set_caption('boby V0.24')
turn ="White"

#Black Pieces
bB = pygame.image.load("images/bB.png")
bK = pygame.image.load("images/bK.png")
bN = pygame.image.load("images/bN.png")
bp = pygame.image.load("images/bp.png")
bQ = pygame.image.load("images/bQ.png")
bR = pygame.image.load("images/bR.png")
#White Pieces
wB = pygame.image.load("images/wB.png")
wK = pygame.image.load("images/wK.png")
wN = pygame.image.load("images/wN.png")
wp = pygame.image.load("images/wp.png")
wQ = pygame.image.load("images/wQ.png")
wR = pygame.image.load("images/wR.png")

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
button_rects = []
piecearr = []
turnCount = 0
gameLog = []

#Game Classes
class Space():
    def __init__(self):
        self.occupied = False
        self.active = False
        self.place = None

class piece():
    def __init__(self):
        self.type = None
        self.val = 0
        self.symbol = None
        self.alive = True
        self.color = None
        self.moved = False
        self.nummoves = 0
        self.boardInd = None
        self.lastMove = None
        self.passant = False
        self.moves = []

    def movegen(self):
        self.moves=[]
        ind = self.boardInd
        if self.type == "ROOK":
            self.moves.extend(self.slides(ind))
        elif self.type == "KNIGHT":
            if board[ind].place[0] == "h":
                self.moves.extend([ind-17, ind-10, ind+6, ind+15])
            elif board[ind].place[0] == "a":
                self.moves.extend([ind-15, ind-6, ind+10, ind+17])
            else:
                self.moves.extend([ind-17, ind-15, ind-10, ind-6, ind+6, ind+10, ind+15, ind+17])
            self.moves = [move for move in self.moves if 0 <= move <= 63]
            self.moves = [move for move in self.moves if ((move%8)-(ind%8)<=2)]
        elif self.type == "BISHOP":
            self.moves.extend(self.diags(ind))
        elif self.type == "QUEEN":
            self.moves.extend(self.slides(ind))
            self.moves.extend(self.diags(ind))
        elif self.type == "KING":
            if ind == "KCastle" or ind == "QCastle":
                if turn == "White":
                    if ind == "Kcastle":
                        self.moves.append(62)
                    elif ind == "Qcastle":
                        self.moves.append(58)
                else:
                    if ind == "Kcastle":
                        self.moves.append(6)
                    elif ind == "Qcastle":
                        self.moves.append(2)
            
            if board[ind].place[0] == "h":
                self.moves.extend([ind-9, ind-8, ind-1, ind+7, ind+8])
            elif board[ind].place[0] == "a":
                self.moves.extend([ind-8, ind-7, ind+1, ind+8, ind+9])
            else: self.moves = ["KCastle", "QCastle", ind-9, ind-8, ind-7, ind-1, ind+1, ind+7, ind+8, ind+9]
        elif self.type == "PAWN":
            if self.color == "White":
                if board[ind].place[1] > 2:
                    self.moved = True
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
                    if Lpassant and (Lpassant.type == "PAWN" and Lpassant.nummoves == 1 and Lpassant.passant == True):
                        self.moves.append(ind-9)
                    if Rpassant and (Rpassant.type == "PAWN" and Rpassant.nummoves == 1 and Rpassant.passant == True):
                        self.moves.append(ind-7)
                    
                    
            elif self.color == "Black":
                if board[ind].place[1] < 7:
                    
                    self.moved = True
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
                    if Lpassant and Lpassant.type == "PAWN" and Lpassant.passant:
                        self.moves.append(ind+7)
                    if Rpassant and Rpassant.type == "PAWN" and Rpassant.passant:
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

                elif piece and piece.color != self.color:
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

    def check(self, turn):
        OppMoves = []
        king = next((piece for piece in piecearr if piece.type == "KING" and piece.color == turn), None)
        for piece in piecearr:
            if piece.color != turn:
                OppMoves.extend(piece.legals(piece.movegen()))
            if king.boardInd in OppMoves:
                return True
        return False

    def EvalMoves(self):
        moves = []
        evals = []
        piece = self
        if not turn:
            return moves
        for move in piece.legals(piece.movegen()):
            if move == ("Kcastle" or "Qcastle") and (not piece.moved):
                moves.append(move)
            elif move == ("Kcastle" or "Qcastle") and (not piece.moved):
                continue
            else:
                ind = piece.boardInd
                moved = piece.moved
                piece.moved = True
                piece.boardInd = move
                if not piece.check(turn):
                    moves.append(move)
                    # key = f"{piece.symbol}{space.place[0]}{space.place[1]}"
                    # val = #eval()
                    # evals.append((key, val))
                    # print((key, val))
                piece.moved = moved
                piece.boardInd = ind
        # print(moves)
        return moves

#Functions Block

#Eval
"""
Piece Activity;
Mobility: Number of moves available to a piece
Centraliztion: Either find the distance from middle rows and columns, or the distance from the current boardInd to inds 27, 28, 35, 36
Attack Potential: Number of pieces attacked by a piece (This will likely involve implementing ray tracing)
Pawn Attacks: Number of pawns attacked by a piece, more potential pawn attacks/captures is more active
Coordination: (This is a stretch rn, but may be useful. Assesses how pieces support each other, and control key squares)
"""
def evals():
    tot = sum(piece.val for piece in piecearr)
    if 0 < abs(tot) < .0001:
        tot -= tot
    AttPot = pieceattack()
    PawnAtt = pawncount()
    Wleft, Bleft = fractionofpieces()
    eval =( .55 * tot) + (.3 * AttPot) + (.1 * PawnAtt) + (.05 * Wleft) - (.05 * Bleft)
    return eval

def minimax(depth, alpha, beta, isMaximizingPlayer):
    pieces = [piece for piece in piecearr if piece.color == turn]

    if depth == 0 or any(piece.check(turn) for piece in pieces):
        ev = evals()
        return ev, None, None

    if isMaximizingPlayer:
        maxEval = float('-inf')
        bestMove = None
        bestPiece = None
        for piece in pieces:
            for move in piece.EvalMoves():
                # Store the original state
                ind = piece.boardInd
                moved = piece.moved
                # Try the move
                piece.boardInd = move
                piece.moved = True

                eval, _, _ = minimax(depth - 1, alpha, beta, False)
                # Restore the original state
                piece.boardInd = ind
                piece.moved = moved

                if eval > maxEval:
                    maxEval = eval
                    bestMove = move
                    bestPiece = piece

                alpha = max(alpha, eval)
                if beta <= alpha:
                    break  # Beta cut-off
        return maxEval, bestMove, board[bestMove].place

    else:
        minEval = float('inf')
        bestMove = None
        bestPiece = None
        for piece in pieces:
            for move in piece.EvalMoves():
                # Store the original state
                ind = piece.boardInd
                moved = piece.moved
                # Try the move
                piece.boardInd = move
                piece.moved = True

                eval, _, _ = minimax(depth - 1, alpha, beta, True)
                # Restore the original state
                piece.boardInd = ind
                piece.moved = moved

                if eval < minEval:
                    minEval = eval
                    bestMove = move
                    bestPiece = piece

                beta = min(beta, eval)
                if beta <= alpha:
                    break  # Alpha cut-off
        return minEval, bestMove, f"{bestPiece.type}, {board[bestMove].place}"


def getAllMoves():
    allMoves = []
    for piece in piecearr:
        if piece.color == turn:
            allMoves.append(piece, piece.EvalMoves())
    return allMoves

#EVAL HELPER FUNCTIONS

def fractionofpieces():
    Wpieces = [piece for piece in piecearr if piece.color == "White"]
    Bpieces = [piece for piece in piecearr if piece.color == "Black"]
    Wleft = len(Wpieces)/16
    Bleft = len(Bpieces)/-16
    return Wleft, Bleft

def EnemyTrace():
    WtraceArr, BtraceArr, Wmovearr, BmoveArr = [], [], [], []

    Wfiltered = [piece for piece in piecearr if piece.color == "White"]
    for piece in Wfiltered:
        Wmovearr.extend(piece.EvalMoves())

    for move in Wmovearr:
        piece2 = next((piece for piece in piecearr if piece.boardInd == move), None)
        if piece2:
            strng = f"{piece2.type}, {piece2.boardInd}, {piece2.val}"
            if piece2.color != piece.color and strng not in WtraceArr:
                WtraceArr.append(strng)

    Bfiltered = [piece for piece in piecearr if piece.color == "Black"]
    for piece in Bfiltered:
        BmoveArr.extend(piece.EvalMoves())

    for move in BmoveArr:
        piece2 = next((piece for piece in piecearr if piece.boardInd == move), None)
        if piece2:
            strng = f"{piece2.type}, {piece2.boardInd}, {piece2.val}"
            if piece2.color != piece.color and strng not in BtraceArr:
                BtraceArr.append(strng)

    return WtraceArr, BtraceArr
    #track pieces in the path of the peices of each side for evaluation purposes

def pieceattack():
    Warr,Barr = EnemyTrace()
    white = len(Warr)
    black = -1 * len(Barr)
    # print(f"White Attacks: {white}, Black Attacks: {black}") #Debugging
    return white + black

def pawncount():
    Warr,Barr = EnemyTrace()
    # print(f"White trace: {Warr}, Black trace: {Barr}") #Debugging
    Wpawncount, Bpawncount = 0, 0
    for item in Warr:
        if "PAWN" in item:
            Wpawncount += 1
    # print(f"White Pawn attacks:{Wpawncount}") #Debugging
    for item in Barr:
        if "PAWN" in item:
            Bpawncount -= 1
    # print(f"Black Pawn attacks:{abs(Bpawncount)}") #Debugging

    fincount = Wpawncount + Bpawncount

    return fincount

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
            WKing.symbol = char
            WKing.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WKing)
            sqr+=1
        elif char == "k":
            BKing = piece()
            BKing.color = "Black"
            BKing.type = "KING"
            BKing.symbol = char
            BKing.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BKing)
            sqr+=1
        elif char == "Q":
            WQueen = piece()
            WQueen.color = "White"
            WQueen.type = "QUEEN"
            WQueen.symbol = char
            WQueen.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WQueen)
            sqr+=1
        elif char == "q":
            BQueen = piece()
            BQueen.color = "Black"
            BQueen.type = "QUEEN"
            BQueen.symbol = char
            BQueen.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BQueen)
            sqr+=1
        elif char == "B":
            WBishop = piece()
            WBishop.color = "White"
            WBishop.type = "BISHOP"
            WBishop.symbol = char
            WBishop.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WBishop)
            sqr+=1
        elif char == "b":
            BBishop = piece()
            BBishop.color = "Black"
            BBishop.type = "BISHOP"
            BBishop.symbol = char
            BBishop.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BBishop)
            sqr+=1
        elif char == "N":
            WKnight = piece()
            WKnight.color = "White"
            WKnight.type = "KNIGHT"
            WKnight.symbol = char
            WKnight.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WKnight)
            sqr+=1
        elif char == "n":
            BKnight = piece()
            BKnight.color = "Black"
            BKnight.type = "KNIGHT"
            BKnight.symbol = char
            BKnight.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BKnight)
            sqr+=1
        elif char == "R":
            WRook = piece()
            WRook.color = "White"
            WRook.type = "ROOK"
            WRook.symbol = char
            WRook.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WRook)
            sqr+=1
        elif char == "r":
            BRook = piece()
            BRook.color = "Black"
            BRook.type = "ROOK"
            BRook.symbol = char
            BRook.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(BRook)
            sqr+=1
        elif char == "P":
            WPawn = piece()
            WPawn.color = "White"
            WPawn.type = "PAWN"
            WPawn.symbol = char
            WPawn.boardInd = sqr
            board[sqr].occupied = True
            piecearr.append(WPawn)
            sqr+=1
        elif char == "p":
            BPawn = piece()
            BPawn.color = "Black"
            BPawn.type = "PAWN"
            BPawn.symbol = char
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
                if m.color == "White":
                    m.val = 3.1
                else: m.val = -3.1
            elif m.type == "KNIGHT":
                if m.color == "White":
                    m.val = 3.0
                else: m.val = -3.0
            elif m.type == "ROOK":
                if m.color == "White":
                    m.val = 5.0
                else: m.val = -5.0
            elif m.type == "QUEEN":
                if m.color == "White":
                    m.val = 9.0
                else: m.val = -9.0
            elif m.type == "PAWN":
                if m.color == "White":
                    m.val = 1.0
                else:  m.val = -1.0
            elif m.type == "KING":
                m.val = 0

#Draws board
def draw():
    global turn, screen
    board_size = (800, 800)
    screen = pygame.display.set_mode(board_size)
    dark = (42, 34, 38)
    light = (22,24,20)
    for j in range(8):  # draw board
        for i in range(8):
            if (i + j) % 2 == 0:
                color = dark
            else:
                color = light
            button = pygame.draw.rect(screen, color, (i * 100, j * 100, 100, 100))
            button_rects.append(button)

    for j in range(8):  # draw pieces
        for i in range(8):
            pos = i + j * 8  # Calculate the position correctly
            if board[pos].occupied:
                piece = next((piece for piece in piecearr if piece.boardInd == pos), None)
                if piece.color == "Black":
                    screen.blit(bPiecesDict[piece.type], (i * 100 + 20, j * 100 + 20))
                else:
                    screen.blit(wPiecesDict[piece.type], (i * 100 + 20, j * 100 + 20))

    pygame.display.flip()

#Game loop and maintainence functions
def go(screen):
    global turn, turnCount, gameLog
    running = True
    dragging=False
    bestMove = None
    clock = pygame.time.Clock()
    #clean up this while loop
    while running:
        initial_alpha = float('-inf')
        initial_beta = float('inf')
        for event in pygame.event.get():
            passantanble = [piece for piece in piecearr if piece.passant == True]
            if event.type == pygame.MOUSEBUTTONDOWN:
                print(f"bestMove: {bestMove}")
                x, y = event.pos
                for inda, button in enumerate(button_rects):
                    if button.collidepoint(x, y):
                        print(f"Button {inda} pressed")
                        orig,piece = takein(inda)
                        ind = inda

                if not orig.occupied or piece.color != turn: #If you click on an empty square or the wrong color
                    drawit(screen)
                    continue
                
                #get king piece and use to check if in check
                king = next((piece for piece in piecearr if piece.type == "KING" and piece.color == turn), None)
                if king.check(turn): #Check evaluation
                    print("IN CHECK")
                    allmoves = []
                    for p in piecearr:
                        if p.color == turn:
                            allmoves += piece.EvalMoves()
                    if allmoves: print("NOT MATE")

                else: #If king is not in check
                    dragging=True
                    if piece:
                        moves = piece.EvalMoves()

                    if piece.color=="White":
                        img = wPiecesDict[piece.type]
                    else:
                        img = bPiecesDict[piece.type]
                

            if event.type== pygame.MOUSEBUTTONUP:
                king = next((piece for piece in piecearr if piece.type == "KING" and piece.color == turn), None)
                dragging = False
                x, y = event.pos
                for indb, button in enumerate(button_rects):
                    if button.collidepoint(x, y):
                        # print(f"Button {indb} pressed , occupied: {board[indb].occupied}")
                        fin,piece2 = takein(indb)
                        ind2 = indb
                if not piece or fin == orig or turn != piece.color: #If no piece or same piece or wrong color
                    drawit(screen)
                
                elif piece2 and (piece2.color == piece.color): #If piece of same color
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
                            if turn == "White":
                                gameLog.append("O-O")
                            else: gameLog[turnCount] = gameLog[turnCount] + " O-O"
                            
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
                            if turn == "White":
                                gameLog.append("O-O-O")
                            else: gameLog[turnCount] = gameLog[turnCount] + " O-O-O"
                    
                    bestMove = None
                    if turn == "Black":
                        turn = "White"
                        Threefold(gameLog[turnCount])
                        turnCount += 1
                        bestMove = minimax(2, initial_alpha, initial_beta, True)
                        print(f"Best Move: {bestMove}")
                    else:
                        turn = "Black"
                        Threefold(gameLog[turnCount])
                        bestMove = minimax(2, initial_alpha, initial_beta, False)
                    drawit(screen)
                    

                elif ind2 in moves: #Check if move is legal
                    if piece.type == "PAWN" and abs(ind2-ind) == 16: #En Passant Block
                        piece.passant = True
                        orig.occupied = False
                        orig.active = False
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True
                        if turn == "White":
                            gameLog.append(f"{board[ind2].place[0]}{board[ind2].place[1]}")
                        else:
                            gameLog[turnCount] = gameLog[turnCount] + f" {board[ind2].place[0]}{board[ind2].place[1]}"


                    elif piece.type == "PAWN" and (abs(ind2-ind) == 9):
                        if turn == "White" and board[ind2+8].occupied and board[ind].place[1]==5: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2+8), None)
                            if piece2.color == "Black" and piece2.passant:
                                orig.occupied = False
                                orig.active = False
                                piece.boardInd = ind2
                                piece.moved = True
                                piece.nummoves += 1
                                fin.occupied = True
                                fin.active = True
                                board[ind2+8].occupied = False
                                if turn == "White":
                                    gameLog.append(f" {board[ind].place[0]}x{board[ind2].place[0]}")
                                else:
                                    gameLog[turnCount] = gameLog[turnCount] + f" {board[ind].place[0]}x{board[ind2].place[0]}"
                                # print(gameLog[turnCount]) #Debugging

                        elif turn == "Black" and board[ind2-8].occupied and board[ind].place[1]==4: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2-8), None)
                            if piece2.color == "White" and piece2.passant:
                                orig.occupied = False
                                orig.active = False
                                piece.boardInd = ind2
                                piece.moved = True
                                piece.nummoves += 1
                                fin.occupied = True
                                fin.active = True
                                board[ind2-8].occupied = False
                                if turn == "White":
                                    gameLog.append(f"{board[ind].place[0]}x{board[ind2].place[0]}")
                                else:
                                    gameLog[turnCount] = gameLog[turnCount] + f" {board[ind].place[0]}x{board[ind2].place[0]}"
                                # print(gameLog[turnCount]) #Debugging
                            
                        else:
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True
                            if turn == "White":
                                gameLog.append(f"{board[ind].place[0]}x{board[ind2].place[0]}")
                            else:
                                gameLog[turnCount] = gameLog[turnCount] + f" {board[ind].place[0]}x{board[ind2].place[0]}"
                            # print(gameLog[turnCount]) #Debugging

                    elif piece.type == "PAWN" and (abs(ind2-ind) == 7): #Grand En Passant block
                        if turn == "White" and board[ind2+8].occupied and board[ind].place[1]==5: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2+8), None)
                            if piece2.color == "Black" and piece2.passant:
                                orig.occupied = False
                                orig.active = False
                                piece.boardInd = ind2
                                piece.moved = True
                                piece.nummoves += 1
                                fin.occupied = True
                                fin.active = True
                                board[ind2+8].occupied = False
                                if turn == "White":
                                    gameLog.append(f" {board[ind].place[0]}x{board[ind2].place[0]}")
                                else:
                                    gameLog[turnCount] = gameLog[turnCount] + f" {board[ind].place[0]}x{board[ind2].place[0]}"
                                print(f" {board[ind].place[0]}x{board[ind2].place[0]}")
                        elif turn == "Black" and board[ind2-8].occupied and board[ind].place[1]==4: #En Passant block
                            piece2 = next((piece for piece in piecearr if piece.boardInd == ind2-8), None)
                            if piece2.color == "White" and piece2.passant:
                                orig.occupied = False
                                orig.active = False
                                piece.boardInd = ind2
                                piece.moved = True
                                piece.nummoves += 1
                                fin.occupied = True
                                fin.active = True
                                board[ind2-8].occupied = False
                                if turn == "White":
                                    gameLog.append(f" {board[ind].place[0]}x{board[ind2].place[0]}")
                                else:
                                    gameLog[turnCount] = gameLog[turnCount] + f" {board[ind].place[0]}x{board[ind2].place[0]}"
                                print(f" {board[ind].place[0]}x{board[ind2].place[0]}")
                        else:
                            orig.occupied = False
                            orig.active = False
                            piece.boardInd = ind2
                            piece.moved = True
                            piece.nummoves += 1
                            fin.occupied = True
                            fin.active = True
                            if turn == "White":
                                gameLog.append(f"{board[ind].place[0]}x{board[ind2].place[0]}")
                            else:
                                gameLog[turnCount] = gameLog[turnCount] + f" {board[ind].place[0]}x{board[ind2].place[0]}"
                            # print(gameLog[turnCount]) #Debugging

                    elif turn == "White" and piece.type == "PAWN" and board[ind2].place[1] == 8: #Promotion block
                        orig.occupied = False
                        orig.active = False
                        piece.type = "QUEEN"
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True
                        if turn == "White":
                            gameLog.append(f"{board[ind2].place[0]}=Q")
                        else:
                            gameLog[turnCount] = gameLog[turnCount] + f" {board[ind2].place[0]}=Q"
                        # print(gameLog[turnCount]) #Debugging

                    elif turn == "Black" and piece.type == "PAWN" and board[ind2].place[1] == 1: #Promotion block
                        orig.occupied = False
                        orig.active = False
                        piece.type = "QUEEN"
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True
                        if turn == "White":
                            gameLog.append(f"{board[ind2].place[0]}=Q")
                        else:
                            gameLog[turnCount] = gameLog[turnCount] + f" {board[ind2].place[0]}=Q"
                        print(f" {board[ind2].place[0]}=Q")    

                    else:
                        orig.occupied = False
                        orig.active = False
                        piece.boardInd = ind2
                        piece.moved = True
                        piece.nummoves += 1
                        fin.occupied = True
                        fin.active = True
                        if piece.type == "PAWN":
                            if turn =="White":
                                gameLog.append(f"{board[ind2].place[0]}{board[ind2].place[1]}")
                            else: 
                                gameLog[turnCount] = gameLog[turnCount] + f" {board[ind2].place[0]}{board[ind2].place[1]}"
                            # print(gameLog[turnCount]) #Debugging
                        elif piece2:
                            if turn =="White":
                                gameLog.append(f"{piece.symbol}x{board[ind2].place[0]}{board[ind2].place[1]}")
                            else: 
                                gameLog[turnCount] = gameLog[turnCount] + f" {piece.symbol}x{board[ind2].place[0]}{board[ind2].place[1]}"
                            # print(gameLog[turnCount]) #Debugging
                        else:
                            if turn =="White":
                                gameLog.append(f"{piece.symbol}{board[ind2].place[0]}{board[ind2].place[1]}")
                            else: 
                                gameLog[turnCount] = gameLog[turnCount] + f" {piece.symbol}{board[ind2].place[0]}{board[ind2].place[1]}"
                            # print(gameLog[turnCount]) #Debugging

                    if piece2:
                        # print(f"Piece Array Length {len(piecearr)}") #Debugging
                        piece2.alive = False 
                        piece2.boardInd = -1
                        piecearr.remove(piece2)
                        # print(f"Piece Array Length {len(piecearr)}") #Debugging


                    bestMove = None
                    if turn == "Black":
                        turn = "White"
                        Threefold(gameLog[turnCount])
                        turnCount += 1
                        # bestMove = minimax(2, initial_alpha, initial_beta, True)
                        print(f"Best Move: {bestMove}")
                    else:
                        turn = "Black"
                        Threefold(gameLog[turnCount])
                        bestMove = minimax(2, initial_alpha, initial_beta, False)
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
    dark = (42, 34, 38)
    light = (22,24,20)
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

def Threefold(move):
    global turn, screen
    occurences = 0
    for element in gameLog:
        if element == move:
            occurences += 1
    if occurences == 3:
        print("Draw by Threefold Repetition")
        turn = None
        drawit(screen)


def takein(ind): #unused, but good to show process.
    spot = board[ind]
    piece = next((piece for piece in piecearr if piece.boardInd == ind), None)

    return spot,piece

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