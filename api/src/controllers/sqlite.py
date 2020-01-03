from peewee import *

from config import Config

class SQLite(object):
    def __init__(self):
        config = Config()
        self.db = SqliteDatabase(config.database)
    
    def get_db(self):
        return self.db
