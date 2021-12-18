## Quiz Game

Create a program that will read a quiz in a CSV file. The quiz is presented 
to a user keeping track of how many questions they get correct and wrong. 
Regardless of whether the answer is correct or wrong the next question should be 
asked immediately afterwards. The CSV file should default to problems.csv, 
but the user should be able to customize the filename via a flag. 
The CSV file format (question,answer):

    5+5,10
    7+3,10
    1+1,2

The quiz will have single word/number answer. At the end of the quiz the program 
should output the number of questions correct and how many questions there were. 
Questions given invalid answers are considered incorrect.

Add to program a timer. The default time limit should be 30 seconds, and also be 
customizable via a flag. The quiz should stop as soon as the time has exceeded.
