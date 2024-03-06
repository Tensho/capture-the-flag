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
