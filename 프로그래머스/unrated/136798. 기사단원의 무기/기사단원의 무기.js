function solution(number, limit, power) {
    let weightCount = 0;
    
    for(let knightNum = 1; knightNum <= number; knightNum++){
        let count = 0;
        for(let i = 1; i ** 2 <= knightNum; i++){
            if(i ** 2 === knightNum) count ++;
            else if(knightNum % i === 0) count += 2;
        };
        count > limit ? weightCount += power : weightCount += count;
    }
    return weightCount;
}