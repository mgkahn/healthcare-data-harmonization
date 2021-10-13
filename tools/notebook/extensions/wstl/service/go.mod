module github.com/GoogleCloudPlatform/healthcare-data-harmonization/tools/notebook/extensions/wstl/service

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

require github.com/google/fhir/go v0.7.0
replace github.com/google/fhir/go v0.7.0 => /fhir-go