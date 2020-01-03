from flask import Blueprint
from config import Config

posts = Blueprint('posts', __name__)

@posts.route('/')
def index():
    return 'hoge'