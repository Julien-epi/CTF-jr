CAPTURE THE FLAG - API

goal : register your username as a result on the specified path

A hint every 20mib
Good luck !

Entrypoint : 10.49.122.144

PATH    METHOD     BODY
/ping   GET 
/signup POST       {"User": string} 
/check  POST       {"User": string}

/getUserSecret  POST  {"User": string}

/getUserLevel   POST  {"User": string, "Secret" : ?}

/getUserPoints  POST   {"User": string, "Secret": ?}

/iNeedAHint     POST    ?
/enterChallenge POST    ?
/submitSolution POST    ?