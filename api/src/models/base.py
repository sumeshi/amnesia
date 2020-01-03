from controllers.sqlite import SQLite

from peewee import *

class BaseModel(Model):
    class Meta:
        database = SQLite().get_db()