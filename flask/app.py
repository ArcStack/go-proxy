import json

from flask import Blueprint, Flask, request
from flask_cors import CORS

app = Flask(__name__)

cors = CORS(app, resources={r"/*": {"origins": "http://localhost:3333"}})
api = Blueprint("api", __name__)


@api.route("/")
def api_hello_world():
    return "Hello World!"


@api.route("/reverse-to-flask")
def api_reverse_to_flask():
    return "Hello from flask!"


app.register_blueprint(api, url_prefix="/api")
if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, port=5000, threaded=True)
