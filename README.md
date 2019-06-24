# go-twilio
Quick setup to respond to a motion event on a raspberry pi, upload the image to S3, and then send a text using twilio


## Usage
`binary {{sid}} {{token}} {{from}} {{to}} {{path to file}}`

This assumes you have the proper AWS access setup, and right now the bucket is hardcoded, so you'll need to change that.

## Epilogue
PRs welcome, this was a quick hit, so I'm sure something isn't right.
