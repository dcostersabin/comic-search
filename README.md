# Comic Search

Cli tool to search offline based on keyword.

Comic-search is a command-line tool that allows you to search for XKCD comics based on keywords. It uses the XKCD JSON API to retrieve data about the comics, including their title, alt text, and URL.



## Environment Variables

To run this project, you will need to add the following environment variables to your .env file. You can take reference from example.env file

```
POSTGRES_USER="<USER>"

POSTGRES_PASSWORD="<PASSWORD>"

POSTGRES_DATABASE="<DATABASE>"

POSTGRES_HOST="<HOST>"

POSTGRES_PORT="<PORT>"

```

## Run Locally

Clone the project

```bash
https://github.com/dcostersabin/comic-search
```

Go to the project directory

```bash
cd comic-search
```

Install dependencies

```bash
go build ./xkcd.go
```

Fetch Data First Before Running Any Search
```bash
./xkcd get -start 990 -end 1000                                                                                                                       130 ↵
Written Comic With Id 990
Written Comic With Id 991
Written Comic With Id 992
Written Comic With Id 993
Written Comic With Id 994
Written Comic With Id 995
Written Comic With Id 996
Written Comic With Id 997
Written Comic With Id 998
Written Comic With Id 1000
```

Search Comic Using Keyword

```bash
./xkcd search -keyword Testing | jq
{
"month": "9",
"num": 632,
"link": "",
"year": "2009",
"news": "",
"safe_title": "Suspicion",
"transcript": "[[A man is sitting at a computer, typing.]]Man: I've loved our online chats these past few months, Lisa.Computer: Me too. I really like you, Rob.[[The man continues to type.]]Man: It's just... now and then you mention products you like, and... I worry.Computer: What? Honey...[[The man types.]]Man: Before this goes any further, I think we should go get tested. You know, together.Computer: You don't trust me?Man: I just want to be sure.[[A web browser is open.]]VK Couples TestingTest ID: 21871138Waiting...Partner connected.((A pair of CAPTCHA images))[You] Library[Partner] KittensMan: Okay, mine says \"library\". Yours?Computer: I... uh...Man: Oh god.Computer: I'm more than a spambot! Our love was real!Man: Goodbye, Lisa.{{Title text: Fine, walk away.  I'm gonna go cry into a pint of Ben&Jerry's Brownie Batter(tm) ice cream [link], then take out my frustration on a variety of great flash games from PopCap Games(r) [link].}}",
"alt": "Fine, walk away.  I'm gonna go cry into a pint of Ben&Jerry's Brownie Batter(tm) ice cream [link], then take out my frustration on a variety of great flash games from PopCap Games(r) [link].",
"img": "https://imgs.xkcd.com/comics/suspicion.png",
"title": "Suspicion",
"day": "4"
}
```

