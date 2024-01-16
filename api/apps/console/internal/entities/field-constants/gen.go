package field_constants

//go:generate go run github.com/kloudlite/api/cmd/struct-json-path --pkg github.com/kloudlite/api/apps/console/internal/entities --common-path "metadata" --common-path "apiVersion" --common-path "kind" --common-path "status" --common-path "syncStatus" --common-path "lastUpdatedBy" --common-path "createdBy" --common-path "displayName" --common-path "creationTime" --common-path "updateTime" --common-path "id" --common-path "recordVersion" --common-path "markedForDeletion"  --out-file ./generated_constants.go --banner "package field_constants" --ignore-nesting "time.Time"
