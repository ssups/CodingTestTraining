// function solution(ingredient, count = 0) {
//     const splited = ingredient.join('').split('1231');
//     count += splited.length - 1;
//     return splited.length > 1 ? solution(splited, count) : count;   
// }

// function solution(ingredient){
//     let count = 0;
//     let keep = true;
    
//     while(keep){
//         ingredient = ingredient.join('').split('1231');
//         count += ingredient.length - 1;
//         keep = ingredient.length > 1 ? true : false;
//     } 
    
//     return count;
// }

// function solution(ingredient, count = 0) {
//     const start = ingredient.join("").indexOf('1231');
//     if(start >= 0){
//         ingredient.splice(start, 4);
//         count ++;
//         return solution(ingredient, count);
//     }
//     else{
//         return count;
//     }
// }

function solution(ingredients){
    const stack = [];
    let count = 0;
    
    ingredients.forEach(ingredient=>{
        stack.push(ingredient);
        if(stack.length >= 4 && stack.slice(-4).join('') === '1231'){
            stack.splice(-4);
            count ++;
        }
    })
    
    return count;
}