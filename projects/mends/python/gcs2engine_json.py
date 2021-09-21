# Move data from GCS to mapping_engine

from google.cloud import storage
from smart_open import open

bucket_name = 'mgk_mends_fhir'
project = 'sandbox-omop'
dataset_id = 'hcls_harmonization_demo'
omop_tables_to_read = ["provider"]

storage_client = storage.Client()
bucket = storage_client.bucket(bucket_name)
if not bucket.exists():
    raise ValueError(
        "The bucket does not exist. Please check the provided bucket name.")

for tbl in omop_tables_to_read:
    blob_filter= "{}/{}-".format(dataset_id,tbl)
    blobs=bucket.list_blobs(prefix=blob_filter)
    for blob in blobs:
        blob_uri = "gs://{}/{}".format(bucket.name,blob.name)
        with open(blob_uri) as fin:
            for line in fin:
                print(line)

# TODO: Shove each line to mapping engine




content = blob.download_as_string()

    if blob and blob.content_encoding:
      content = content.decode(blob.content_encoding)

    # check if the returned content is a json
    try:
      try:
        result = json.loads(content)
      except TypeError:
        result = json.loads(content.decode("UTF-8"))
    except json.JSONDecodeError:
      print(
          "The loaded content is not a valid JSON. Please check the source bucket and blob."
      )
      raise
    if dest_file_name:
      with open(dest_file_name, "w") as dest:
        dest.write(content)
        return "The message was written to {} successfully.".format(
            dest_file_name)

    return JSON(result)

for 
