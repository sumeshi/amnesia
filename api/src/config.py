from pathlib import Path


class Config(object):
    def __init__(self):
        self.project_root = Path(__file__).parent
        self.assets_directory = self.project_root / Path('static')
        self.database = self.assets_directory / Path('database.sqlite3')