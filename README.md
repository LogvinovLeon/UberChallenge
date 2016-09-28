# UberChallenge
My email sender for the UBER challenge.
https://github.com/uber/coding-challenge-tools/blob/master/coding_challenge.md

## TL;DR
Just go and try it! [www.uberchallenge.email](www.uberchallenge.email)
The API: [https://api.uberchallenge.email/email](https://api.uberchallenge.email/email) (I would describe it later)

##Problem description:
This service allows you to send emails using different email providers.
For now it supports [Mailgun](https://mailgun.com) and [Sendgrid](https://sendgrid.com). 
 
The best and the only feature of this service is that it falls back to the secondary provider if first is down.

* Mailgun has 99.99% SLA.
* Sendgrid does not have an SLA.

## Scope
The solution focuses on the back-end, but also has a front-end part 
which allows to play with the service and send some emails.

## Architecture

Back-end and front-end are completely separated.
They have nothing in common. 

I just placed them into a single repo so it would be easier for you to review.

### Back-end
Back-end is a Go app.
I have absolutely no experience with Go and started learning it after I decided to write this project.

#### Why Go?

* Go is designed to be fast, safe and reliable.
* It's statically typed.
* It is also designed for web.
* This service is IO-heavy. It spends most of the time communicating with the providers.
* Go runs every request in it's own goroutine, so if one request get's stuck or slow or fails it would not affect other users.

The whole back-end contains of a single endpoint which allows you to (surprise!) send emails.
It you want to test it you can execute this command from console:
```
curl -X POST -H "Content-Type: application/json" \
-d='{"to":"logvinov.leon@gmail.com", \
"subject": "Lorem", \
"body": "ipsum"}' \
https://api.uberchallenge.email/email -i
```

I wanted to keep it as simple as possible, that's why I'm not using any frameworks.
You don't need them for such a simple app.
It also has `preferred_provider` parameter which is not required.
You can set it to `sendgrid` if you're brave. By default we start sending with `mailgun`.

####Back-end is deployed on Heroku and this gives me a couple of benefits

* Monitoring & metrics
* Logging
* Continuous deployment (github integration and easy rollbacks)
* Scalability (the servers are completely stateless, so I can just add more instances with a single command)

If this night this service would become EXTREMELY popular I CAN (but wouldn't because of the costs) handle the load.
I would receive an email from mailgun, that I would reach the quota soon.
The only two actions that I need to do is to change the mailgun plan and scale the number of heroku instances.

If I would have more time and motivation I would use some of the auto-scaling solutions, but it seemed to be an overkill.

#### Security
The service checks the correctness of the client input and uses https.

I do not store any tokens/API keys in the repo.

In order to run the service you need to set those environment variables:
```
SENDGRID_API_KEY
MAILGUN_DOMAIN
MAILGUN_API_KEY
MAILGUN_API_KEY
PORT (heroku sets it)
```

And execute:
`go run web.go`

If you're planning to use only one provider you don't need to set the variables for the other one.


### Font-end
Front-end is in [docs](docs) directory.
It's just a couple of static files, so you can run it on your local machine.
I'm serving them using github pages connected to a custom domain.
GitHub Pages sites have a bandwidth limit of 100GB or 100,000 requests per month.
If my front-end would become EXTREMELY popular at night I would receive a polite email from GitHub (they guarantee it) and would setup a CDN in front of GH pages.
GH pages is already a CDN, so it also allows we to serve my page faster to the users in different locations.
#### Security
I do not have any confidential or sensitive information on those pages, that's why I don't use https for Front-end.
(Also because GH Pages does not support https for custom domains yet)
If it would be critical I can setup a CDN in front of it/instead of it and have https.
#### Speed
I did not use any heavy css frameworks like bootstrap.
I used [milligram](https://milligram.github.io/) which is only 2kB gzipped.

# Other projects I'm proud about

* [Interpreter of my own functional language written in Haskell](https://github.com/LogvinovLeon/MIML)
* [Algorithmic trading bot for Jane Street hackathon](https://github.com/LogvinovLeon/eth1)
* Awesome work on Profilers and JS packaging that I've done at Quora, but cannot describe here because of an NDA.

# Links:

* [My CV](https://docs.google.com/document/d/1wxfYc1kwj5c51uXhXfhgFe0ZjelIctBDrPelmXH3VBA/edit) (It has all other links inside)
* [My Github](https://github.com/LogvinovLeon)


