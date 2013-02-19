go-twitter-notifications
========================

Go application that pushes twitter notifications matching a set of patterns to your Pebble.

Installation
============

Clone the repository, and then do:

    go get

That will download the dependencies. After that, you need to modify create config.json. 
You should copy ```config.sample.json``` and modify the configuration fields with your information.

After that, just install it: 

    go install

And if you have $GOPATH/bin in your $PATH, just execute it:

    $ go-twitter-push &

Configuring match patterns
==========================

You can configure an arbitrary number of MatchPatterns. Each one of them has a `From` and `MatchText`.
"*" is a *wildcard*. It will match any user or text.

    "Patterns": [
        {"From": "justinbieber", "MatchPattern": "new song"},
        {"From": "*", "MatchPattern": "hunter2"}
        /* etc */
    ]
 
