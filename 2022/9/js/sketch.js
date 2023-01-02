let cols, rows;
let w = 40;
let grid = [];
let input;
let start;

let index = 0;

function preload() {
  data = loadStrings('/data/test.txt');
}

function setup() {
  createCanvas(440, 440);
  cols = floor(width/w);
  rows = floor(height/w);

  for (let j = 0; j < rows; j++) {
    for (let i = 0; i < cols; i++) {
      grid.push(new Cell(i, j));
    }
  }

  head = grid[floor(cols/2)*cols + floor(rows/2)];
}

function draw() {
  background(51);

  head.visited = true;

  let direction = data[index].slice(0,1);
  let scalar = parseInt(data[index].slice(1));
  
  visited = 0;
  for (let i = 0; i < grid.length; i++) {
    grid[i].show();
    if (grid[i].visited) {
      visited++;
    }
  }


  console.log('Visited count:', visited, ', Current Direction:', direction, scalar);
  noLoop();
}
