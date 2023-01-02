class Cell {
    constructor(i, j, letter) {
        this.i = i;
        this.j = j;
        this.letter = letter;
        this.height = this.anum(letter)
        this.neighbors = [];
        this.walls = [true, true, true, true];
        this.visited = false;
    }

    index(i, j) {
        if (i < 0 || j < 0 || i > cols-1 || j > rows-1) {
            return -1
        }
        return i + j * cols;
    }

    anum(l) {
        var anum={
            a: 0,  b: 1,  c: 2,  d: 3,  e: 4,  f: 5,  g: 6,  h: 7,  i: 8,  j: 9,
            k: 10, l: 11, m: 12, n: 13, o: 14, p: 15, q: 16, r: 17, s: 18, t: 19, 
            u: 20, v: 21, w: 22, x: 23, y: 24, z: 25, S:-1,  E:26
        }

        return anum[l]
    }

    checkNeighbors() {
        var neighbors = [];

        var current = grid[this.index(this.i, this.j)];

        var top = grid[this.index(this.i, this.j - 1)];
        var right = grid[this.index(this.i + 1, this.j)];
        var bottom = grid[this.index(this.i, this.j + 1)];
        var left = grid[this.index(this.i - 1, this.j)];
  
        if (top && !top.visited) {
            if (top.height - current.height == 1 || top.height - current.height == 0) {
                current.walls[0] = false;
                top.walls[2] = false;
                neighbors.push(top);
            }
        }

        if (right && !right.visited) {
            if (right.height - current.height == 1 || right.height - current.height == 0) {
                current.walls[1] = false;
                right.walls[3] = false;
                neighbors.push(right)
            }
        }

        if (bottom && !bottom.visited) {
            if (bottom.height - current.height == 1 || bottom.height - current.height == 0) {
                current.walls[2] = false;
                bottom.walls[0] = false;
                neighbors.push(bottom);
            }
        }

        if (left && !left.visited) {
            if (left.height - current.height == 1 || left.height - current.height == 0) {
                current.walls[3] = false;
                left.walls[1] = false;
                neighbors.push(left);
            }
        }
        return neighbors
    }

    highlight() {
        var x = this.i * w;
        var y = this.j * w;
        noStroke();
        fill(0, 255, 0, 100);
        rect(x, y, w, w);
    }

    show() {
        var x = this.i * w;
        var y = this.j * w;

        stroke(255);
        // top
        if (this.walls[0]) {
            line(x, y, x + w, y);
        }
        // right
        if (this.walls[1]) {
            line(x + w, y, x + w, y + w);
        }
        // bottom
        if (this.walls[2]) { 
            line(x + w, y + w, x, y + w);
        }
        //left
        if (this.walls[3]) {
            line(x, y+  w, x, y)
        }
        if (this.visited) {
            noStroke();
            fill(255, 0, 255, 100);
            rect(x, y, w, w);
          }
        
        textAlign(CENTER, CENTER)
        noStroke();
        text(this.letter, x+w/2, y+w/2)
        fill(255);
    }
}