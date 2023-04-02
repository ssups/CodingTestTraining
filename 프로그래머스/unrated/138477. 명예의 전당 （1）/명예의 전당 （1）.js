function solution(k, score) {
    const ranker = [];
    const result = []
    
    score.forEach(el=>{
        ranker.push(el);
        ranker.sort((a, b)=>b - a);
        const length = ranker.length
        if(ranker.length > k) ranker.length = k;
        result.push(ranker[ranker.length - 1]);
    });
    
    return result;
}