#!/bin/sh

function blog() {
  gcloud functions deploy BlogHTTP --runtime go113 --trigger-http --entry-point=BlogHandler --region=asia-northeast1
}

$1
