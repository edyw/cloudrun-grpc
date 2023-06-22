# gRPC Service to Service on Cloud Run

This is a simple example of [gRPC](https://grpc.io/) Service to Service using Go and Node JS. Services are designed for GCP [Cloud Run](https://cloud.google.com/run) deployment, leveraging Google Cloud authentication and IAM authorization, private networking design for Serverless.

You can find the full article [here](https://medium.com/google-cloud/grpc-service-to-service-on-cloud-run-132f3963d183).

## Quick start with gRPC service to service call, open to Internet, SSL enabled, http/2, no authentication 


1. Clone this repository and download dependencies:
    ```
    git clone https://github.com/edyw/cloudrun-grpc.git
    ```
2. Create GCP Project, enable services
    ```
    PROJECT_ID=your-project-id

    gcloud projects create ${PROJECT_ID}
    gcloud config set project ${PROJECT_ID}

    gcloud services enable run.googleapis.com
    gcloud services enable artifactregistry.googleapis.com
    ```
3. Create Artifact Registry repository
    ```
    REGION=your-region
    REPO=your-artifact-registry-repository-name

    gcloud artifacts repositories create ${REPO} \
        --repository-format=docker \
        --location=${REGION}
    ```
4. Deploy go-contact gRPC service
    ```
    docker build . -f ./go-contact/Dockerfile --tag ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/go-contact

    docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/go-contact

    gcloud run deploy go-contact --image=${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/go-contact \
        --region=${REGION} \
        --project=${PROJECT_ID} \
        --allow-unauthenticated \
        --use-http2
    ```
5. Replace go-api/main.go contactServerHost variable with go-contact Cloud Run Service URL

6. Deploy go-api REST service
    ```
    docker build . -f ./go-api/Dockerfile --tag ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/go-api

    docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/go-api

    gcloud run deploy go-api --image=${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO}/go-api \
        --region=${REGION} \
        --project=${PROJECT_ID} \
        --allow-unauthenticated \
        --port=8081
    ```
7. End to end test
    ```
    GO_API_CLOUD_RUN_SERVICE=cloud-run-go-api-service-url

    curl ${GO_API_CLOUD_RUN_SERVICE}/contact/555-110022
    ```
    You should have output: Mike F%

