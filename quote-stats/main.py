from flask import Flask,render_template 
import requests
from os import getenv
import requests

app = Flask(__name__)
url = getenv("QUOTES_API_URL")
"""
response object
{
 {
  "quote_count": 0,
  "last_quote": {
    "title": "someone also said",
    "body": "alfred kind of cool"
  }
}
"""

def getstats(url):
    response = requests.get(url).json()
    return response 



@app.route('/')
def stats():
    response = getstats(url)
    quote_count = response['quote_count']
    last_quote = response['last_quote']
    return render_template('index.html',quote_count=quote_count,last_quote=last_quote)

