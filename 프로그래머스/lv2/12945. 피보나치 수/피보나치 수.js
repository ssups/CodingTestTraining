function solution(n) {
    const fibonachi = [0 ,1];
    for(let i = 2; i <= n ; i++){
        const next = ((fibonachi[i-1]) + (fibonachi[i-2]))% 1234567;
        fibonachi.push(next);
    }
    return fibonachi[n] ;
}