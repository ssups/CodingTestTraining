// function solution(n, m, section) {
//     return m === 1 ? n : paint(section, m);
// }

// function paint(sections, meter, count = 0){
//     const lastSection = sections[0] + meter - 1;
//     const newSections = sections.filter(section => section > lastSection);
//     count ++;
//     return newSections.length === 0 ?  count : paint(newSections, meter, count);
// }

function solution(n, m, section) {
    let lastPainted = 0;
    let count = 0;
    
    section.forEach(tile => {
        if(tile > lastPainted){
            lastPainted = tile + m - 1;
            count ++;
        }; 
    });
    
    return count;
}