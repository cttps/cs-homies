const yourAge = document.getElementById("yourAge");
const theirAge = document.getElementById("theirAge");
const result = document.getElementById("result");
const rationale = document.getElementById("rationale");
let verdict;



function calculate(){

    yourAgeNum = Number(yourAge.value);
    theirAgeNum = Number(theirAge.value);


    minAgeForYou = yourAgeNum / 2 + 7;
    minAgeForThem = theirAgeNum / 2 + 7;

    if (yourAge.value === '' || theirAge.value === ''){
        result.textContent = "You forgot to put in the values my friend.";
    }

    else if (yourAgeNum < 14 && theirAgeNum < 14){
        result.textContent = "You are both under 14, so the rule doesn't really work for those ages.";
        rationale.textContent = ""
    }



    else if(theirAgeNum < minAgeForYou){
        result.textContent = "The FBI has been alerted.";
        rationale.textContent = `The rule states the minimum age you can date is ${minAgeForYou}.
        It's time to end the relationship.`
    } 

    else if(yourAgeNum < minAgeForThem){
        result.textContent = "You are a victim.";
        rationale.textContent = `The rule states the minimum age they can date is ${minAgeForThem}.
        It's time to end the relationship.`
    } 

    else{
        result.textContent = "Your relationship is acceptable!";
        rationale.textContent = `The rule states the minimum age you can date is ${minAgeForYou} 
        and the minimum age they can date is ${minAgeForThem}.
        It all looks good from here.`
    }

}