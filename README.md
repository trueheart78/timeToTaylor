# TimeToTaylor

Waiting for a concert kinda stinks.

![bored - taylor swift.gif](https://dl.dropboxusercontent.com/s/spnflcr84vl993i/bored+-+taylor+swift.gif)

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

![taylor - looks good.gif](https://dl.dropboxusercontent.com/s/emq30wgry1usuhs/taylor+-+looks+good.gif)
