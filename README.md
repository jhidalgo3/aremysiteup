# AreMySiteUp

Original idea [__Bulldog__](https://github.com/pioz/bulldog) thanks @pioz

AreMySiteUp is a monitor process that checks for you a list of URLs and warns you by email using [Mailgun](https://www.mailgun.com/) if one of them returns a http code that is not 200.

## Installation

Install it yourself

    $ go get github.com/jhidalgo3/aremysiteup
    $ cd $GOPATH/srcgithub.com/jhidalgo3/aremysiteup
    $ go build
    $ ./aremysiteup -v

## Usage

AreMySiteUp loads configuration from $PWD/configs/aremysiteup.yml:


```yaml
# After checking the entire list of URLs sleep for these seconds.
sleep: 60
# After checking the entire list of URLs and at least a check fail sleep for
# these seconds. Usually this time is greater to not warn you continuously.
sleepWithError: 600
# Http request timeout. If the timeout is reached the check is to be considered
# as failed.
timeout: 10

# Disables logs.
quiet: false

# Mailgun configuration
mailgun:
  domain: domain.com
  apiKey: API key - you can get this value from the Mailgun admin interface.
  publicApiKey: is the public API key - you can get this value from the Mailgun admin interface. 


# When a check fails send an email on this email address. If is empty the email
# alert is disabled.
to: user@mail.com

from: "no-reply <no-reply@domain.com>"

# Comma-separated list of URLs to check.
urls: 
  - http://google.com
  - http://www.twitter.coma
```

To unleash run the follow command:

    $ aremysiteup

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/jhidalgo3/aremysiteup

## License

The package is available as open source under the terms of the [MIT License](https://github.com/jhidalgo3/aremysiteup/blob/master/LICENSE).