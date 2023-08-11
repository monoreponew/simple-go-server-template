external/gcloud/gcloud services enable run.googleapis.com --project=$GCP_PROJECT

external/crane/crane auth login gcr.io -u oauth2accesstoken -p $CLOUDSDK_AUTH_ACCESS_TOKEN

external/crane/crane push server/go_image.tar gcr.io/$GCP_PROJECT/go_image:latest

external/gcloud/gcloud run deploy test-$RANDOM --image=gcr.io/$GCP_PROJECT/go_image:latest --port=3333 --region=us-west1 --project=$GCP_PROJECT --allow-unauthenticated