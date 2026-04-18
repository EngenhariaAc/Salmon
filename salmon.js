class Salmon {
  constructor(name, weight) {
    this.name = name;
    this.weight = weight;
  }

  swim() {
    console.log(`${this.name} is swimming upstream! 🐟`);
  }
}

module.exports = Salmon;
