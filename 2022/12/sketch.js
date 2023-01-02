var cols, rows;
var w = 20;
var grid = [];

var current;
var stack = [];

function preload() {
  input = loadStrings("data/test.txt");
}

function setup() {
  rows = input.length;
  cols = input[0].length;

  createCanvas(cols*w, rows*w);

  // debugging
  // frameRate(10);

  for (let j = 0; j < rows; j++) {
    for (let i = 0; i < cols; i++) {
      let cell = new Cell (i, j, input[j][i])
      grid.push(cell);
      if (cell.letter == "S") {
        current = cell
      }
    }
  }
}



function draw() {
  background(51);
  for (let i = 0; i < grid.length; i++) {
    grid[i].checkNeighbors();
    grid[i].show();
  }
}