class Cell {
    constructor(i,j) {
        this.i = i;
        this.j = j;
        this.visited = false;        
    }

    move(direction) {
        switch (direction) {
            case 'U':
                this.j++;
                break;
            case 'R':
                this.i++
                break;
            case 'D':
                this.j--;
                break;
            case 'L':
                this.i--;
                break;
            default:
                break;
        }
    }

    show() {
        var x = this.i * w;
        var y = this.j * w;
        stroke(255);
        noFill();
        rect(x, y, w, w);

        if (this.visited) {
            noStroke();
            fill(255, 0, 255, 100)
            rect(x, y, w, w);
        }
    }
}