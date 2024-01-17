# FullCycle temperature system
Get temperature by CEP

# 🧪 Testing in Local

# 🚀 Deploy to GCP Cloud Run
- Set these variables on your favourite shell:
  - `ENV_PROJECT_ID`: This is the projectID that you are using, for example in my case is: `generic-316622`
  - `ENV_ARG`: this is the environment variable, you can use either `prod` or `dev`.
- You need to log to gcloud using `gcloud auth login`
- Then, run `gcloud config set project ${ENV_PROJECT_ID}`
- You need to execute `deploy.sh` using somethin like this, (you mush change `ENV_PROJECT_ID`)
  ```shell
  ENV_PROJECT_ID=generic-316622 ENV_ARG=prod ./deploy.sh
  ```
- In my case, I have this URL to test: `https://temperature-uz6otgp6yq-uc.a.run.app`
