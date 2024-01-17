#!/bin/bash

gcloud builds submit --config cloudbuild.yaml --substitutions=_PROJECT_ID="${ENV_PROJECT_ID}",_ENV_ARG="${ENV_ARG}"

gcloud run deploy temperature \
    --image "gcr.io/${ENV_PROJECT_ID}/temperature" \
    --region us-central1 \
    --project "${ENV_PROJECT_ID}" \
    --allow-unauthenticated \
    --platform managed
