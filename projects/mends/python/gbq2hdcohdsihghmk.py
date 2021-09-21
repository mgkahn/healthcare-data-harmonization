# Move data from GBQ to GCS as JSON

#!/usr/bin/env python3

from google.cloud import bigquery

client = bigquery.Client()
bucket_name = 'hdcohdsihghmk'
project = 'hdcohdsihghmk'
dataset_id = 'n3c_mends'
omop_tables_to_extract = ["condition_occurrence","drug_exposure","person","provider","visit_occurrence"]

for tbl in omop_tables_to_extract:
    blob_name = "mgk_mends_fhir/{}-*.json".format(tbl)
    destination_uri = "gs://{}/{}".format(bucket_name, blob_name)
    dataset_ref = bigquery.DatasetReference(project, dataset_id)
    table_ref = dataset_ref.table(tbl)
    job_config = bigquery.job.ExtractJobConfig()
    job_config.destination_format = bigquery.DestinationFormat.NEWLINE_DELIMITED_JSON
    
    extract_job = client.extract_table(
        table_ref,
        destination_uri,
        job_config=job_config, 
        location="US", # Location must match that of the source table.
        )  # API request
        
    extract_job.result()  # Waits for job to complete.
