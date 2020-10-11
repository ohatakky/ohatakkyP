#!/bin/sh

function blog_function() {
  gcloud functions deploy BlogHTTP --runtime go113 \
  --trigger-http \
  --entry-point=BlogHandler \
  --region=asia-northeast1 \
  --timeout=540 \
  --env-vars-file .env.yaml
  # --ingress-settings=internal-only \
}

function blog_scheduler() {
  gcloud scheduler jobs create BlogScheduler http \
  --schedule="*/30 * * * *" \
  --time-zone="Asia/Tokyo" \
  --uri=${BLOG_FUNCTION_URI} \
  --oidc-service-account-email=${BLOG_FUNCTION_SERVICE_ACCOUNT}
}

$1
