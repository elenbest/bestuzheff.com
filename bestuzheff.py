from flask import Flask
app = Flask(__name__)

index_file = open('index.html', 'r')


@app.route('/')
def hello_world():
    return index_file.read()


if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0.0.0.0", port=8080)
