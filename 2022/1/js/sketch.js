let elves = [];
let input;
let index = 0;

let binWidth;
let maxCal = 0;

function preload() {
  input = loadStrings("test.txt");
}
function setup() {
  createCanvas(400, 400);
  background(155);

  let calories = [];
  for (let i = 0; i < input.length; i++) {
    let value = parseFloat(input[i]);

    if (value) {
      calories.push(value)
    } else {
      elf = new Bar(index, calories);
      elves.push(elf);
      if (sum(calories) > maxCal) {
        maxCal = sum(calories)
      }
      calories= [];
      index++;
    }
  }
  binWidth = width/elves.length
}

function draw() {
  for (let i = 0; i < elves.length; i++) {
    elves[i].display();
  }
  
}

function sum(vals) {
  let sum = 0;
  for (let i = 0; i < vals.length; i++) {
    sum+=vals[i];
  }
  return sum;
}

function total(vals) {
  let total = 0;
  for (let i = 0; i < vals.length; i++) {
    total = total + vals[i];
  }
  return total;
}