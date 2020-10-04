#!/bin/sh

function blog_function() {
  gcloud functions deploy BlogHTTP --runtime go113 \
  --trigger-http \
  --entry-point=BlogHandler \
  --region=asia-northeast1 \
  --env-vars-file .env.yaml
  # --ingress-settings=internal-only \
}


# https://cloud.google.com/scheduler/docs/http-target-auth?hl=ja#using-gcloud
function blog_scheduler() {
  gcloud scheduler jobs create http --schedule="every 15 mins" --uri=${BLOG_FUNCTION_URI} --oidc-service-account-email=${BLOG_FUNCTION_SERVICE_ACCOUNT} --time-zone="Asia/Tokyo"
}

$1
