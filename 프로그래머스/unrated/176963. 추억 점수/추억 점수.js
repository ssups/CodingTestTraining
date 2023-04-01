function solution(name, yearning, photo) {
    const mapping = name.reduce((acc,cur,ind)=>{
        acc[cur] = yearning[ind];
        return acc;
    },{});
    
    const result = photo.map(arr=>{
        return arr.reduce((acc,name)=>{
            return acc + (mapping[name] || 0);
        },0)
    });
    return result;
}