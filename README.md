# TimeToTaylor

Waiting for a concert kinda stinks.

![bored - taylor swift.gif][bored-taylor-swift]

Writing Go code that works with AWS Lambda makes it go faster.

## Installing It

Download the relevant file from the [releases list](https://github.com/trueheart78/timeToTaylor/releases) and extract it.

Now you know!

## Development

### Locally

Running it locally:

```
$ go install timeToTaylor.go

$ timeToTaylor
Taylor Swift is in Columbus, Ohio on Saturday, July 7th, 2018 at 7pm

[ 66 days, 3 hours, 36 minutes, 43 seconds ]
```

### AWS Lambda

Generate a ready-made AWS Lambda archive.

```
$ bin/build-lambda

Creating release dir...
  adding: out/aws-lambda (deflated 65%)

AWS Lambda Handler Info: 'out/aws-lambda'
```

## Profit

![taylor - looks good.gif][looks-good-taylor-swift]


[bored-taylor-swift]: assets/bored-taylor-swift.gif
[looks-good-taylor-swift]: assets/looks-good-taylor-swift.gif
