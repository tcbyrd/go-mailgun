### Mailgun Contact Form
- A very simple app written in Go to handle a contact form submission and setup the message for processing through Mailgun. 
- Returns 404 for anything that's not a POST request to the `/mailer` function.
- Reads Mailgun configuration from environment variables
- Confifigured for Heroku deployment using the `heroku/go` buildpack.
