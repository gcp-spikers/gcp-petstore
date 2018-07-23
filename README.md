# gcp-petstore

## Starting up

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

#### Environment Vars:

- `PROJECT_ID` should be your GCP Project Id

#### Example invocation

        PROJECT_ID=[PROJECT_ID] \
        REPO_URL="https://github.com/crccheck/docker-hello-world.git" \
            container-builder-local \
            --dryrun=false \
            --substitutions _REPO_URL=$REPO_URL,_SHOULD_RUN_IMAGE=1 \
            .

  

<!--
    REFERENCES
-->

[gcb-docker-quickstart]: https://cloud.google.com/container-builder/docs/quickstart-docker