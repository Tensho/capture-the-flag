# Capture The Flag (CTF) - Hackathon 2024

Find out 3 flags.

### Prerequisites

    $ brew bundle
    $ go mod init livelink/ctf
    $ go get
    $ go build

## Development

    $ go build
    $ ./main
    $ curl -I http://localhost:8080

## Deployment (GCP)

    $ gcloud config set core/project com-livelinklabs-ctf
    $ gcloud auth login
    $ gcloud functions deploy ctf --gen2 \
                                  --runtime go121 \
                                  --trigger-http \
                                  --allow-unauthenticated \ 
                                  --max-instances 1 \
                                  --set-secrets=FLAG_1=projects/765124037022/secrets/FLAG_1/versions/1,FLAG_2=projects/765124037022/secrets/FLAG_2/versions/1,FLAG_3=projects/765124037022/secrets/FLAG_3/versions/1

## Misc

### Get available functions runtimes

    $ gcloud functions runtimes list
