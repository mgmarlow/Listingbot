## Listingbot

A Slack bot that posts new apartment listings from craigslist to a channel.

## Usage
Create a `settings.json` file and drop it in the same directory as the executable. The contents of the file should look like:

```
{
    "slackToken": "my-slack-token",
    "city": "my-craigslist-city-url",
    "daysPast": <how many days to search>,
    "price": <max-pricepoint>
}
```

For example:

```
{
    "slackToken": "my-slack-token",
    "city": "santabarbara",
    "daysPast": 3,
    "price": 950
}
```

You can run the `run.sh` shell script to compile and run the executable. This process assumes you have go installed.
