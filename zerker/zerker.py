import fastapi
import uvicorn
import pprint
import berserk

app = fastapi.FastAPI()

@app.get("/tablebase")
async def tableBase(fenStr, api_tok):
    print(fenStr)
    try:
        session= berserk.TokenSession(api_tok)
        client = berserk.clients.Tablebase(session=session, tablebase_url="tablebase.lichess.ovh")

        res = client.look_up(position=fenStr)
        pprint.pprint(res)
    except Exception:
        return "NOT IMPLEMENTED"
    return "NOT IMPLEMENTED"

@app.get("/eval")
async def eval(fenStr, api_tok)->float|tuple[float,str]:
    try:
        print(fenStr)
        session = berserk.TokenSession(api_tok)
        client = berserk.clients.Analysis(session=session, base_url="https://lichess.org")

        res = client.get_cloud_evaluation(fen=fenStr, num_variations=0)
        score = float(res["pvs"][0]["cp"])
        pprint.pprint(f"score: {score}")
        return score    
    except Exception:
        return (404.0 , f"No cloud evaluation available for position {fenStr}")

def seekGames(api_tok):
    games = {internal: 0,
             seek: 0,
             challenge: 0}
    while games < 3:
        session = berserk.TokenSession(api_tok)
        board = berserk.clients.Board(session=session, base_url="https://lichess.org")
        board.stream_incoming_events()
        games += 1


# app.on_event("startup")(seekGames)

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=7000)

    
