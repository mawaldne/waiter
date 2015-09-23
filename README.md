# Waiter

"A person who waits for a time, event, or opportunity."

A simple http endpoint that sleeps for a certain amount of time. Mainly
used in webhook testing. 

If you somehow manager to figure out how to build this go project:

PORT=3000 ./main

curl -X POST https://localhost:3000/wait   -> wait 1 second  
curl -X POST https://localhost:3000/wait?seconds=5   -> wait 5 second  

Mostly used with just heroku. 
