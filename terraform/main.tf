terraform {
  backend gcs {
    bucket = "livelink-terraform"
    prefix = "capture-the-flag.tfstate"
  }
}

provider "google" {
  project = "com-livelinklabs-ctf"
  region  = "europe-west2"
}

resource "google_project_service" "cloudfunctions" {
  service            = "cloudfunctions.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "cloudbuild" {
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "run" {
  service            = "run.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "artifactregistry" {
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "logging" {
  service            = "logging.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "secretmanager" {
  service            = "secretmanager.googleapis.com"
  disable_on_destroy = false
}

# TODO: Add default service account IAM permission to read secrets

resource "random_string" "flag" {
  count = 3

  length  = 64
  special = false

  keepers = {
    secret = google_secret_manager_secret.flag[count.index].secret_id
  }
}

resource "google_secret_manager_secret" "flag" {
  count = 3

  secret_id = "FLAG_${count.index + 1}"

  replication {
    auto {}
  }

  depends_on = [google_project_service.secretmanager]
}

resource "google_secret_manager_secret_version" "flag" {
  count = 3

  secret      = google_secret_manager_secret.flag[count.index].id
  secret_data = random_string.flag[count.index].result
}
