class Bar {
    constructor(i, arr) {
        this.y = i;
        this.arr = arr;
        this.r = random(255);
        this.g = random(255);
        this.b = random(255);
    }

    total(arr) {
        this.w = 0;
        for (let i = 0; i < this.arr.length; i++) {
            this.w += this.arr[i]
        }
    }

    display() {
        fill(this.r, this.g, this.b, 25)
        rect(0, this.y*binWidth, map(total(this.arr), 0, maxCal, 0, width), binWidth)
    }
}