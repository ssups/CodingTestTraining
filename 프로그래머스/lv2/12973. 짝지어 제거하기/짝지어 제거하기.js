function solution(string){
    const stack = [];
    for(const each of string){
        stack.push(each);
        while(stack.length >= 2 && stack[stack.length - 1] === stack[stack.length - 2]){
            stack.splice(-2);
        } 
    }
    return stack.length === 0 ? 1 : 0;
}