# gcp-petstore

## Application options

### Command line arguments

- `-port int` port number (default `8080`)

### Environment variables for settings

- `PETSTORE_FAIL` cause the app to report "Unhealthy" (defaults to healthy)
- `PETSTORE_CRASH` crash the app on first request (defaults to false)
- `PETSTORE_CRASHTIMER_MIN` - `PETSTORE_CRASHTIMER_MAX`
Specifies a range for a random timer which crashes when time is up (defaults to `-1` to deactivate)

## Running locally

The Dockerfile is tailored to using Google Cloud Build, but if you have `go` installed you can try it locally with:

    go run main.go

## Starting up in-cloud

### Attaching to a GCP project

Don't forget to authorise your `gcloud` CLI tool

    gcloud auth login

And set your project to match GCP configuration, where `[PROJECT_ID]` is your GCP project ID

    gcloud config set project [PROJECT_ID]

### Running pipeline

    gcloud container builds submit \
    --tag gcr.io/[PROJECT_ID]/pipeline-product:latest .


### Running pipeline locally

You need `container-builder-local` installed, and will need to provide some environment vars and/or substitutions.

#### Environment Variables for local invocation

- `PROJECT_ID` should be your GCP Project Id
- `REPO_URL` specify an alternate repository (not active)

#### Example invocation

    PROJECT_ID=[PROJECT_ID] \
    REPO_URL="https://github.com/gcp-spikers/gcp-petstore.git" \
        container-builder-local \
        --dryrun=false \
        --substitutions _REPO_URL=$REPO_URL,_SHOULD_RUN_IMAGE=1 \
        .

  

<!--
    REFERENCES
-->

[gcb-docker-quickstart]: https://cloud.google.com/container-builder/docs/quickstart-docker
