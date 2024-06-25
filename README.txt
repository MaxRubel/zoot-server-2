to build: 


$ gcloud auth login

## optional
$ gcloud config set project zoot-server
##

# Build the Docker image
docker build -t zoot-server .

# Tag the image for GCR
docker tag zoot-server gcr.io/zoot-server/zoot-server

# Push the image to GCR
docker push gcr.io/zoot-server/zoot-server

# Deploy to Cloud Run
gcloud run deploy zoot-server --image gcr.io/zoot-server/zoot-server --platform managed