class Rope {
    constructor(i, j) {
        this.i = i;
        this.j = j;
    }

    move(direction) {
        switch (direction) {
            case 'U':
              this.j++;
              break;
            case 'R':
              this.i++;
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
        noFill(52, 235, 120);
        rect(x, y, w, w);
    }
}