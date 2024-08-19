import seedrandom from "seedrandom";

export class RandomManager {
    private static _instance: RandomManager;

    private random: any = null;

    /** 随机数管理单例对象 */
    static get instance(): RandomManager {
        if (this._instance == null) {
            this._instance = new RandomManager();
            this._instance.random = Math.random;
        }
        return this._instance;
    }

    private getRandom(): number {
        return this.random();
    }

    /** 设置随机种子 */
    setSeed(seed: number) {
        this.random = new seedrandom(seed);
    }

    /**
     * 生成指定范围的随机浮点数
     * @param min   最小值
     * @param max   最大值
     * @param type  类型
     */
    getRandomFloat(min: number = 0, max: number = 1): number {
        return this.getRandom() * (max - min) + min;
    }

    /**
     * 生成指定范围的随机整数
     * @param min   最小值
     * @param max   最大值
     * @param type  类型
     * @example
     let min = 1;
     let max = 10;
     // [min,max) 得到一个两数之间的随机整数,这个值不小于min（如果min不是整数的话，得到一个向上取整的 min），并且小于（但不等于）max
     RandomManager.instance.getRandomInt(min, max, 1);

     // [min,max] 得到一个两数之间的随机整数，包括两个数在内,这个值比min大（如果min不是整数，那就不小于比min大的整数），但小于（但不等于）max
     RandomManager.instance.getRandomInt(min, max, 2);

     // (min,max) 得到一个两数之间的随机整数
     RandomManager.instance.getRandomInt(min, max, 3);
     */
    getRandomInt(min: number, max: number, type: number = 2): number {
        min = Math.ceil(min);
        max = Math.floor(max);
        switch (type) {
            case 1: // [min,max) 得到一个两数之间的随机整数,这个值不小于min（如果min不是整数的话，得到一个向上取整的 min），并且小于（但不等于）max  
                return Math.floor(this.getRandom() * (max - min)) + min;
            case 2: // [min,max] 得到一个两数之间的随机整数，包括两个数在内,这个值比min大（如果min不是整数，那就不小于比min大的整数），但小于（但不等于）max
                return Math.floor(this.getRandom() * (max - min + 1)) + min;
            case 3: // (min,max) 得到一个两数之间的随机整数
                return Math.floor(this.getRandom() * (max - min - 1)) + min + 1;
        }
        return 0;
    }

    /**
     * 根据最大值，最小值范围生成随机数数组
     * @param min   最小值
     * @param max   最大值
     * @param n     随机个数
     * @param type  类型
     * @example
     let a = RandomManager.instance.getRandomByMinMaxList(50, 100, 5)
     console.log("随机的数字", a);
     */
    getRandomByMinMaxList(min: number, max: number, n: number): Array<number> {
        let result: Array<number> = [];
        for (let i = 0; i < n; i++) {
            result.push(this.getRandomInt(min, max))
        }
        return result;
    }

    /**
     * 获取数组中随机对象
     * @param objects 对象数组
     * @param n 随机个数
     * @example
     let b = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
     let r = RandomManager.instance.getRandomByObjectList(b, 5);
     console.log("原始的对象", b);
     console.log("随机的对象", r);
     */
    getRandomByObjectList<T>(objects: Array<T>, n: number): Array<T> {
        let temp: Array<T> = objects.slice();
        let result: Array<T> = [];
        for (let i = 0; i < n; i++) {
            let index = this.getRandomInt(0, objects.length, n);
            result.push(temp.splice(index, 1)[0]);
        }
        return result;
    }

    /**
     * 定和随机分配
     * @param n     随机数量
     * @param sum   随机元素合
     * @example
     let c = RandomManager.instance.getRandomBySumList(5, -100);
     console.log("定和随机分配", c);
     */
    getRandomBySumList(n: number, sum: number): number[] {
        let residue = sum;
        let value = 0;
        let result: Array<number> = [];
        for (let i = 0; i < n; i++) {
            value = this.getRandomInt(0, residue, 3);
            if (i == n - 1) {
                value = residue;
            } else {
                residue -= value;
            }
            result.push(value);
        }
        return result;
    }

    randomWord(hasChar: boolean, min: number, max: number) {
        let str = "", range = min;
        let arr = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
        if (hasChar) {
            arr = [
                '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
                'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A',
                'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
                'U', 'V', 'W', 'X', 'Y', 'Z'
            ]
        }
        for (let i = 0; i < range; i++) {
            let pos = Math.round(Math.random() * (arr.length - 1));
            str += arr[pos];
        }
        return str;
    }
}