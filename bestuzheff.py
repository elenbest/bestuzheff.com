from flask import Flask
from waitress import serve
app = Flask(__name__)


@app.route('/')
def hello_world():
    index_file = open('index.html', 'r')
    return index_file.read()


if __name__ == '__main__':
    serve(app, host='0.0.0.0', port=80)
