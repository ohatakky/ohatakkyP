#!/bin/sh

function blog_function() {
  gcloud functions deploy BlogHTTP --runtime go113 \
  --trigger-http \
  --entry-point=BlogHandler \
  --region=asia-northeast1 \
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

# function trending_function() {
#   gcloud functions deploy TrendingHTTP --runtime go113 \
#   --trigger-http \
#   --entry-point=TrendingHandler \
#   --region=asia-northeast1 \
#   --env-vars-file .env.yaml
#   # --ingress-settings=internal-only \
# }

# function trending_scheduler() {
#   gcloud scheduler jobs create TrendingScheduler http \
#   --schedule="0 10 * * *" \
#   --time-zone="Asia/Tokyo" \
#   --uri=${TRENDING_FUNCTION_URI} \
#   --oidc-service-account-email=${TRENDING_FUNCTION_SERVICE_ACCOUNT}
# }

$1
