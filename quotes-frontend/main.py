from distutils.log import error
from flask import Flask 
from os import getenv
from flask.templating import render_template 
from typing import List
import requests
api_url = getenv("QUOTES_API_URL")

app = Flask(__name__)


def GetQuotes(url) -> List[str]:
    response:List[str] = requests.get(url).json()
    return response

@app.route("/")
def index():
    quotes = GetQuotes(api_url)
    return render_template("index.html",quotes=quotes)
