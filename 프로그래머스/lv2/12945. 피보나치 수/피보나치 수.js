function solution(n) {
    const fibonachi = [0 ,1];
    for(let i = 2; i <= n ; i++){
        const next = BigInt(fibonachi[i-1]) + BigInt(fibonachi[i-2]);
        fibonachi.push(next);
    }
    return fibonachi[n] % 1234567n;
}