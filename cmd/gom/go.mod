module main

go 1.14

require GoMetadata/cmd/gom v1.0.0
replace GoMetadata/cmd/gom => ./
require internal/metadata v1.0.0
replace internal/metadata => ../../internal/metadata
