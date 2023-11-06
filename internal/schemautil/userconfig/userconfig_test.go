//go:build userconfig
// +build userconfig

package userconfig

import (
	"fmt"
	"testing"

	"github.com/aiven/go-api-schemas/pkg/dist"
	"github.com/dave/jennifer/jen"
	"github.com/ettle/strcase"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// generateSchema generates Terraform schema from a map representation of the schema.
func generateSchema(schemaName string, schemaMap map[string]any) error {
	schemaNamePlural := fmt.Sprintf("%ss", schemaName)

	file := jen.NewFile("dist")

	file.HeaderComment("Code generated by internal/schemautil/userconfig/userconfig_test.go; DO NOT EDIT.")

	sortedMapKeys := maps.Keys(schemaMap)
	slices.Sort(sortedMapKeys)

	for _, key := range sortedMapKeys {
		value := schemaMap[key]

		keyPascalCase := strcase.ToGoPascal(key)

		valueAsserted, ok := value.(map[string]any)
		if !ok {
			continue
		}

		properties, ok := valueAsserted["properties"].(map[string]any)
		if !ok {
			continue
		}

		functionName := fmt.Sprintf("%s%s", schemaName, keyPascalCase)

		file.Commentf(
			"%s is a generated function returning the schema of the %s %s.", functionName, key, schemaName,
		)

		required := map[string]struct{}{}

		if requiredSlice, ok := valueAsserted["required"].([]any); ok {
			required = SliceToKeyedMap(requiredSlice)
		}

		propertiesMap, err := convertPropertiesToSchemaMap(properties, required)
		if err != nil {
			return err
		}

		file.
			Func().
			Id(functionName).
			Params().
			Id("*schema.Schema").
			Block(
				jen.Id("s").Op(":=").Map(jen.String()).Op("*").Qual(
					SchemaPackage, "Schema",
				).Values(propertiesMap),

				jen.Line(),

				jen.Return(
					jen.Op("&").Qual(SchemaPackage, "Schema").Values(jen.Dict{
						jen.Id("Type"): jen.Qual(SchemaPackage, "TypeList"),
						jen.Id("Description"): jen.Lit(fmt.Sprintf(
							"%s user configurable settings", keyPascalCase,
						)),
						jen.Id("Elem"): jen.Op("&").Qual(SchemaPackage, "Resource").
							Values(jen.Dict{jen.Id("Schema"): jen.Id("s")}),
						jen.Id("Optional"): jen.Lit(true),
						jen.Id("DiffSuppressFunc"): jen.Qual(
							SchemaUtilPackage, "EmptyObjectDiffSuppressFuncSkipArrays",
						).Call(jen.Id("s")),
						jen.Id("MaxItems"): jen.Lit(1),
					}),
				),
			).
			Line()
	}

	if err := file.Save(fmt.Sprintf("dist/%s.go", strcase.ToSnake(schemaNamePlural))); err != nil {
		return err
	}

	return nil
}

// TestMain is the entry point for the user config schema generator.
func TestMain(m *testing.M) {
	serviceTypesMap, err := representationToMap(ServiceTypes, dist.ServiceTypes)
	if err != nil {
		panic(err)
	}

	err = generateSchema("ServiceType", serviceTypesMap)
	if err != nil {
		panic(err)
	}

	integrationTypesMap, err := representationToMap(IntegrationTypes, dist.IntegrationTypes)
	if err != nil {
		panic(err)
	}

	err = generateSchema("IntegrationType", integrationTypesMap)
	if err != nil {
		panic(err)
	}

	integrationEndpointTypesMap, err := representationToMap(IntegrationEndpointTypes, dist.IntegrationEndpointTypes)
	if err != nil {
		panic(err)
	}

	err = generateSchema("IntegrationEndpointType", integrationEndpointTypesMap)
	if err != nil {
		panic(err)
	}
}
