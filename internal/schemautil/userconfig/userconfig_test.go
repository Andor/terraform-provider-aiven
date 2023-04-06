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

// generateSchema is a function that generates Terraform schema via its map representation.
func generateSchema(n string, m map[string]interface{}) error {
	np := fmt.Sprintf("%ss", n)

	f := jen.NewFile("dist")

	f.HeaderComment("Code generated by internal/schemautil/userconfig/userconfig_test.go; DO NOT EDIT.")

	smk := maps.Keys(m)
	slices.Sort(smk)

	for _, k := range smk {
		v := m[k]

		kp := strcase.ToGoPascal(k)

		va, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		pa, ok := va["properties"].(map[string]interface{})
		if !ok {
			continue
		}

		fn := fmt.Sprintf("%s%s", n, kp)

		f.Commentf("%s is a generated function returning the schema of the %s %s.", fn, k, n)

		pm, err := convertPropertiesToSchemaMap(pa, map[string]struct{}{})
		if err != nil {
			return err
		}

		f.
			Func().
			Id(fn).
			Params().
			Id("*schema.Schema").
			Block(
				jen.Id("s").Op(":=").Map(jen.String()).Op("*").Qual(SchemaPackage, "Schema").Values(pm),

				jen.Line(),

				jen.Return(
					jen.Op("&").Qual(SchemaPackage, "Schema").Values(jen.Dict{
						jen.Id("Type"):        jen.Qual(SchemaPackage, "TypeList"),
						jen.Id("Description"): jen.Lit(fmt.Sprintf("%s user configurable settings", kp)),
						jen.Id("Elem"): jen.Op("&").Qual(SchemaPackage, "Resource").Values(jen.Dict{
							jen.Id("Schema"): jen.Id("s"),
						}),
						jen.Id("Optional"): jen.Lit(true),
						jen.Id("DiffSuppressFunc"): jen.
							Qual(SchemaUtilPackage, "EmptyObjectDiffSuppressFuncSkipArrays").Call(jen.Id("s")),
						jen.Id("MaxItems"): jen.Lit(1),
					}),
				),
			).
			Line()
	}

	if err := f.Save(fmt.Sprintf("dist/%s.go", strcase.ToSnake(np))); err != nil {
		return err
	}

	return nil
}

// TestMain is the entry point for the user config schema generator.
func TestMain(m *testing.M) {
	stm, err := representationToMap(ServiceTypes, dist.ServiceTypes)
	if err != nil {
		panic(err)
	}

	err = generateSchema("ServiceType", stm)
	if err != nil {
		panic(err)
	}

	itm, err := representationToMap(IntegrationTypes, dist.IntegrationTypes)
	if err != nil {
		panic(err)
	}

	err = generateSchema("IntegrationType", itm)
	if err != nil {
		panic(err)
	}

	ietm, err := representationToMap(IntegrationEndpointTypes, dist.IntegrationEndpointTypes)
	if err != nil {
		panic(err)
	}

	err = generateSchema("IntegrationEndpointType", ietm)
	if err != nil {
		panic(err)
	}
}
