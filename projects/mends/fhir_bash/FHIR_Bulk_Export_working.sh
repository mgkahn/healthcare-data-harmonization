# FHIR server @ https://healthcare.googleapis.com/v1/projects/sandbox-omop/locations/us-central1/datasets/hcls_qs/fhirStores/synthea_r4:export
# FHIR bulk GCS bucket output @ gs://mgk_mends_fhir/Bulk_out

curl -X POST \
  -H "Authorization: Bearer  $(gcloud auth application-default print-access-token)" \
  -H "Content-Type: application/json; charset=utf-8" \
  --data "{'gcsDestination':{'uriPrefix':'gs://mgk_mends_fhir/Bulk_out'}}" \
  "https://healthcare.googleapis.com/v1/projects/sandbox-omop/locations/us-central1/datasets/hcls_qs/fhirStores/synthea_r4:export"
