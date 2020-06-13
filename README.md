# spanner

Cloud Spanner tutorial

## How to develop

```bash
$ cd cloudrun
$ go run main.go
```

## How to deploy

```bash
$ export PROJECT_ID=<your-gcp-project-id>
$ docker build --target deployment -t asia.gcr.io/$PROJECT_ID/spanner:latest .
$ docker push asia.gcr.io/$PROJECT_ID/spanner:latest
```

## How to run Github Actions

### 1. Create service account

- Cloud Storage Admin
- Cloud Run Invoker

### 2. Register the following secrets by referring to [creating-and-storing-encrypted-secrets](https://help.github.com/ja/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets)

- GCP_PROJECT_ID: <your-gcp-project-id>
- GCP_REGION: <your-gcp-region>（e.g. asia-northeast1）
- GCP_SERVICE_ACCOUNT_EMAIL: <your-gcp-service-account>
- GCP_SERVICE_ACCOUNT_KEY: <your-gcp-service-account-key>
  - You need generate base64 character from key-json file.

```bash
  $ cat key-file.json | base64
```
