class Person(object):
    def __init__(self, person_dict):
        self.id = person_dict.get('id')
        self.name = person_dict.get('name')
        self.surname = person_dict.get('surname')
        self.age = person_dict.get('age')
