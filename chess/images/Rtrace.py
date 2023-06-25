turn ="White"
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
    symbol = None
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
                OppMoves.extend(piece.legals(piece.movegen(piece.boardInd)))
            if king.boardInd in OppMoves:
                return True
        return False
            
    def isLegalMove(self, ind):
        boardcopy = board
        piecearrcopy = piecearr
        finalArr = []
        pieces = [piece for piece in piecearrcopy if piece.color == turn]
        for piece in pieces:
            for move in piece.legals(piece.movegen(piece.boardInd)):
                piece.makeMove(move)
                if not piece.check(ind):
                    finalArr.append(move)
                piece.undoMove(move)
        return finalArr
                

    def makeMove(self, ind):
        pass

    def undoMove(self, ind):
        pass

def trace():
    #track pieces in the path of the peices of each side for evaluation purposes
    pass
