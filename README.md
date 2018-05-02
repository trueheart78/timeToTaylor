# TimeToTaylor

Waiting for a concert kinda stinks.

![bored - taylor swift.gif](https://dl.dropboxusercontent.com/s/spnflcr84vl993i/bored+-+taylor+swift.gif)

Writing Go code that works with AWS Lambda makes it go faster.

## Locally

Running it locally:

```
$ go install timeToTaylor.go

$ timeToTaylor
Taylor Swift is in Columbus, Ohio on Saturday, July 7th, 2018 at 7pm

[ 66 days, 3 hours, 36 minutes, 43 seconds ]
```

## AWS Lambda

Generate a reaady-made AWS Lambda archive.

```
$ bin/build-lambda

Creating release dir...
  adding: out/aws-lambda (deflated 65%)

AWS Lambda Handler Info: 'out/aws-lambda'
```

## Profit

![dancing in the rain - taylor swift.gif](https://dl.dropboxusercontent.com/s/lj1vm953hnc56zn/dancing+in+the+rain+-+taylor+swift.gif)
