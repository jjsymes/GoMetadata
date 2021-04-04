module GoMetadata

go 1.14

require internal/metadata v1.0.0
replace internal/metadata => ./internal/metadata
require cmd/gom v1.0.0
replace cmd/gom => ./cmd/gom
// replace github.com/jjsymes/greetings => ../greetings
// require example.com/greetings v1.1.0
