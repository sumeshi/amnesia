from models.base import BaseModel

from peewee import *

class Post(BaseModel):
    id = AutoField()
    title = CharField()
    