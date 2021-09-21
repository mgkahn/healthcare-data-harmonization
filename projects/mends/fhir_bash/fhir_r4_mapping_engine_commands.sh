# Single resource. Must run in $DH_ROOT/mapping_engine/main directory

#WORKS
go run $DH_ROOT/mapping_engine/main/. \
--input_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/ConditionOccurence.fhir.input.json \
--mapping_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/configurations/main.wstl \
--lib_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/projector_library/ \
--harmonize_code_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/code_harmonization/

# Mupltiple resources. Must run in $DH_ROOT/mapping_engine/main directory
#WORKS
go run $DH_ROOT/mapping_engine/main/. \
--input_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/\*.fhir.input.json \
--mapping_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/configurations/main.wstl \
--lib_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/projector_library/ \
--harmonize_code_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/code_harmonization/


#WORKS -- one ndjson file on local disk
# MENDS configs
go run $DH_ROOT/mapping_engine/main/. \
	--input_file_spec=$MENDS_ROOT/data/condition1.json \
	--mapping_file_spec=$MENDS_ROOT/mapping_configs/omop_fhir_r4/configurations/main.wstl \
	--lib_dir_spec=$MENDS_ROOT/mapping_configs/omop_fhir_r4/projector_library/ \
	--harmonize_code_dir_spec=$MENDS_ROOT/mapping_configs/omop_fhir_r4/code_harmonization/




#WORKS -- writes to local output_dir
go run $DH_ROOT/mapping_engine/main/. \
--input_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/ConditionOccurence.fhir.input.json \
--output_dir=$DH_ROOT/projects/mends/output/ \
--mapping_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/configurations/main.wstl \
--lib_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/projector_library/ \
--harmonize_code_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/code_harmonization/




#DOES NOT WORK
# Pulling input JSON from GCS
pushd $DH_ROOT/mapping_engine/main
go run $DH_ROOT/mapping_engine/main/. \
--input_file_spec=gs://sandbox-omop/mgk_mends_fhir/hcls_harmonization_demo/*.fhir.input.json \
--output_dir=$DH_ROOT/projects/mends/output/ \
--mapping_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/configurations/main.wstl \
--lib_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/projector_library/ \
--harmonize_code_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/code_harmonization/
popd





dlv debug main.go -- --input_file_spec=gs://sandbox-omop/mgk_mends_fhir/hcls_harmonization_demo/*.fhir.input.json --output_dir=$DH_ROOT/projects/mends/output/ --mapping_file_spec=$DH_ROOT/mapping_configs/omop_fhir_r5/configurations/main.wstl --lib_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/projector_library/ --harmonize_code_dir_spec=$DH_ROOT/mapping_configs/omop_fhir_r4/code_harmonization/

