let input;
let index = 0;
let calories = [];

function preload() {
  input = loadStrings("test.txt");
}
function setup() {
  createCanvas(400, 400);
  background(0);
}

function draw() {
  let value = parseFloat(input[index]);

  if (value) {
    calories.push(value);
  } else {
    console.log(calories, total(calories));
    calories = [];
  }


  if (index == input.length) {
    noLoop();
  }

  index++;
}

function total(vals) {
  let total = 0;
  for (let i = 0; i < vals.length; i++) {
    total = total + vals[i];
  }
  return total;
}