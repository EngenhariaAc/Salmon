class Fish:
    def __init__(self, species, habitat):
        self.species = species
        self.habitat = habitat

    def describe(self):
        return f"{self.species} lives in {self.habitat}."
