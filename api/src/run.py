import os

from flask import Flask

from views.posts import posts

os.path.abspath(os.path.join(os.path.dirname( __file__ ), '.', 'src'))

app = Flask(__name__)
app.register_blueprint(posts, url_prefix='/posts')

if __name__ == '__main__':
    app.run(host='0.0.0.0')