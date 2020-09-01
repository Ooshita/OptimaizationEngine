from flask import Flask
from TabuSearch import TabuSearch

app = Flask(__name__)

@app.route('/')
def tabu_search():
    ts = TabuSearch()
    return ts.execution(100000)