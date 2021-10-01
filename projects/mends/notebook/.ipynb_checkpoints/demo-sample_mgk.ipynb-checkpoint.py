get_ipython().run_line_magic("load_ext", " wstl")


get_ipython().run_cell_magic("markdown", "", """
## Load JSON data directly from Google Cloud Storage

The following line magic command loads JSON data from a GCS bucket and 
allows the user to either save the results to a python variable, persist them to disk, or print the results.""")


get_ipython().run_line_magic("load_hl7v2_gcs?", "")


get_ipython().run_line_magic("load_hl7v2_gcs", " --bucket_name mgk_mends_fhir --source_blob_name hcls_harmonization_demo/ConditionOccurence.fhir.input.json")


get_ipython().run_cell_magic("markdown", "", """
## Evaluate whistle mapping language
* allows the user to execute whistle mapping language scripts and to immediately see the result of executing the script
* result can be displayed in output or stored into a python variable.
* inputs and library configs can be read from local filesystem or Google Cloud Storage
* outputs a FHIR transaction bundle""")


get_ipython().run_line_magic("%wstl?", "")


get_ipython().run_cell_magic("markdown", "", """
### Reads from the file system mounted into the docker container.""")


import os
os.chdir('/home/jupyter')
os.getcwd()


get_ipython().run_cell_magic("wstl", " --output resources --input file://./mapping_configs/omop_fhir_r4/ConditionOccurence.fhir.input.json --library_config file://./mapping_configs/omop_fhir_r4/projector_library/*.wstl --code_config file://./mapping_configs/omop_fhir_r4/code_harmonization/*.json", """
$this : Process_ConditionOccurence($root.ConditionOccurence);
    
// Description: Process OMOP (v6.0) ConditionOccurence to construct FHIR R4 Condition
//
// Argument(s):
//   ConditionOccurence: https://github.com/OHDSI/CommonDataModel/blob/v6.0.0/Documentation/CommonDataModel_Wiki_Files/StandardizedClinicalDataTables/CONDITION_OCCURRENCE.md
//
// Output(s):
//   Condition: https://www.hl7.org/fhir/condition.html
//


// post Output_Transaction

""")


print(resources)


#https://storage.cloud.google.com/mgk_mends_fhir/hcls_harmonization_demo/adt_a01.hl7.fhir.input.json

get_ipython().run_line_magic("load_hl7v2_gcs", " --bucket_name mgk_mends_fhir --source_blob_name hcls_harmonization_demo/adt_a01.hl7.fhir.input.json")



get_ipython().run_cell_magic("markdown", "", """
## Evaluate whistle mapping language
* allows the user to execute whistle mapping language scripts and to immediately see the result of executing the script
* result can be displayed in output or stored into a python variable.
* inputs and library configs can be read from local filesystem or Google Cloud Storage
* outputs a FHIR transaction bundle""")


get_ipython().run_line_magic("%wstl?", "")


get_ipython().run_cell_magic("markdown", "", """
### Reads from the file system mounted into the docker container.""")


import os
os.chdir('/home/jupyter')
os.getcwd()


get_ipython().run_cell_magic("wstl", " --output resources --input file://./mapping_configs/hl7v2_fhir_stu3/adt_a01_msh.hl7.fhir.input.json --library_config file://./mapping_configs/hl7v2_fhir_stu3/projector_library/*.wstl", """
$this : Process_ADT_A01($root.ADT_A01);

// Description: Define HL7v2 (version 2.9) to FHIR STU3 transformation
//
// Argument(s):
//   ADT_A01 HL7v2 Message Type (version 2.9)
//   ADT_A01 Segments covered in this mapping configuration are: MSH
//
// Output(s):
//   FHIR STU3 Resources : http://www.hl7.org/fhir/stu3/resourcelist.html
// List of Generated FHIR Resources:
//   Bundle: https://www.hl7.org/fhir/stu3/bundle.html

def Process_ADT_A01(ADT_A01) {
  var MH : MSH_MessageHeader(ADT_A01.MSH);
  out Resources : MH;
  out Resources : HD_Provenance(ADT_A01.MSH.3, MH);
  out Resources : MSH_Bundle(ADT_A01.MSH.10);
}

// post Output_Transaction
""")


get_ipython().run_cell_magic("markdown", "", """
### Reads from Google Cloud Storage""")


get_ipython().run_cell_magic("wstl", " --output resources --input gs://data-harmonization-sample-data/hl7v2_fhir_stu3/adt_a01_msh.hl7.fhir.input.json --library_config gs://data-harmonization-sample-data/hl7v2_fhir_stu3/projector_library/*.wstl", """
$this : Process_ADT_A01($root.ADT_A01);
    
// Description: Define HL7v2 (version 2.9) to FHIR STU3 transformation
//
// Argument(s):
//   ADT_A01 HL7v2 Message Type (version 2.9)
//   ADT_A01 Segments covered in this mapping configuration are: MSH
//
// Output(s):
//   FHIR STU3 Resources : http://www.hl7.org/fhir/stu3/resourcelist.html
// List of Generated FHIR Resources:
//   Bundle: https://www.hl7.org/fhir/stu3/bundle.html

def Process_ADT_A01(ADT_A01) {
  var MH : MSH_MessageHeader(ADT_A01.MSH);
  out Resources : MH;
  out Resources : HD_Provenance(ADT_A01.MSH.3, MH);
  out Resources : MSH_Bundle(ADT_A01.MSH.10);
}""")


get_ipython().run_cell_magic("markdown", "", """
### Mappings for ADT_A01 PID and PD1 segments""")


get_ipython().run_cell_magic("wstl", " --output resources --input file://./mapping_configs/hl7v2_fhir_stu3/adt_a01_pid_pd1.hl7.fhir.input.json --library_config file://./mapping_configs/hl7v2_fhir_stu3/projector_library/*.wstl --code_config file://./mapping_configs/hl7v2_fhir_stu3/code_harmonization/*.json", """
$this : Process_ADT_A01($root.ADT_A01);

// Description: Define HL7v2 (version 2.9) to FHIR STU3 transformation
//
// Argument(s):
//   ADT_A01 HL7v2 Message Type (version 2.9)
//   ADT_A01 Segments covered in this mapping configuration are: PID, PD1,
//
// Output(s):
//   FHIR STU3 Resources : http://www.hl7.org/fhir/stu3/resourcelist.html
// List of Generated FHIR Resources:
//   Account: https://www.hl7.org/fhir/stu3/account.html
//   Organization: https://www.hl7.org/fhir/stu3/organization.html
//   Patient: https://www.hl7.org/fhir/stu3/patient.html
//   Practitioner: https://www.hl7.org/fhir/stu3/practitioner.html
//
def Process_ADT_A01(ADT_A01) {
  var Patient : PID_PD1_Patient(ADT_A01.PID, ADT_A01.PD1);
  out Resources : Patient;
  out Resources : PID_Account(ADT_A01.PID.18, Patient);
}""")


get_ipython().run_cell_magic("markdown", "", """
## Load Data from HL7v2Store
The following line magic command allows you to load data directly from a HL7v2Store and store the results either in a python variable or persist them to disk. 
You will manually need to do the following:
    
* Create an [HL7v2 store](https://cloud.google.com/healthcare/docs/how-tos/hl7v2#creating_an_hl7v2_store) on Google Cloud with 
[schematize parsing](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.hl7V2Stores#schematizedparsingtype) enabled. 
For example, the "schematizedParsingType" attribute set to "SOFT_FAIL" and the "ignoreMinOccurs" attribute set to true.
The schematized parsing options enable the store to generated a structured JSON representatio of HL7v2 messages similar to the the example representations
above. 
* [Add HL7v2 messages](https://cloud.google.com/healthcare/docs/how-tos/hl7v2-messages) to the store.
* Update the magic command in the cell below with the appropiate store details.""")


get_ipython().run_line_magic("load_hl7v2_datastore?", "")


# uncomment, replace all the labels in square brackets, and execute the following cell magic command
# hl7v2_store_messages = get_ipython().run_line_magic("load_hl7v2_datastore", " --project_id [PROJECT_ID] --region [REGION] --dataset_id [DATASET] --hl7v2_store_id [STORE_ID]")


import json
parsed_messages = [json.loads(x['schematizedData']['data']) for x in hl7v2_store_messages.data]

from IPython.display import JSON
parsed_message = parsed_messages[0]
JSON(parsed_message)


get_ipython().run_cell_magic("markdown", "", """
### Reads from python variable previous defined in cell""")


get_ipython().run_cell_magic("wstl", " --output resources --input py://parsed_message --library_config file://./mapping_configs/hl7v2_fhir_stu3/projector_library/*.wstl", """
$this : Process_ORU_R01($root.ORU_R01);
    
// Description: Define HL7v2 (version 2.9) to FHIR STU3 transformation
//
// Argument(s):
//   ORU_R01 HL7v2 Message Type (version 2.9)
//   ORU_R01 Segments covered in this mapping configuration are: MSH
//
// Output(s):
//   FHIR STU3 Resources : http://www.hl7.org/fhir/stu3/resourcelist.html
// List of Generated FHIR Resources:
//   Bundle: https://www.hl7.org/fhir/stu3/bundle.html

def Process_ORU_R01(ORU_R01) {
  var MH : MSH_MessageHeader(ORU_R01.MSH);
  out Resources : MH;
  out Resources : HD_Provenance(ORU_R01.MSH.3, MH);
  out Resources : MSH_Bundle(ORU_R01.MSH.10);
}""")



