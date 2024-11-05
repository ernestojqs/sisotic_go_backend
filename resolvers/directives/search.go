package directives

import (
	"encoding/json"

	"github.com/pjmd89/gogql/lib/gql"
	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Search struct {
	scalars gql.Scalars
}

func NewSearch(scalars gql.Scalars) resolvers.Directive {
	d := &Search{scalars: scalars}
	return d
}
func (o *Search) Invoke(args map[string]interface{}, typeName string, fieldName string) (r resolvers.DataReturn, err definitionError.GQLError) {
	search := bson.D{}
	if args != nil && args["input"] != nil {
		for _, v := range args["input"].([]interface{}) {
			search = append(search, o.parse(v.(map[string]interface{}))...)
		}
	}
	r = search
	return r, err
}
func (o *Search) parse(args map[string]interface{}) (r []bson.E) {
	value := args["value"].([]interface{})
	and := args["and"].([]any)
	not := args["not"].([]any)
	nor := args["nor"].([]any)
	or := args["or"].([]any)
	jsonbody1, _ := json.Marshal(args["geoSpatial"].(map[string]interface{})["geoWithin"])
	geoWithinOperators := &GeoSpatialGeoWithinOperators{}
	json.Unmarshal(jsonbody1, &geoWithinOperators)
	jsonbody2, _ := json.Marshal(args["geoSpatial"].(map[string]interface{})["near"])
	geoSpatialNear := &GeoSpatialNear{}
	json.Unmarshal(jsonbody2, &geoSpatialNear)
	if len(value) > 0 && args["field"] != nil && len(and) == 0 && len(not) == 0 && len(nor) == 0 && len(or) == 0 {
		bsonE := bson.E{}
		bsonE.Key = args["field"].(string)
		bsonM := bson.M{}
		for _, vValue := range value {
			switch vValue.(map[string]any)["operator"].(string) {
			case "in", "nin", "all":
				parseValue := []any{}
				values := vValue.(map[string]any)["values"].([]any)
				for _, vValues := range values {
					parsevValues, _ := o.scalars[vValue.(map[string]any)["kind"].(string)].Set(vValues.(string))
					parseValue = append(parseValue, parsevValues)
				}
				bsonM["$"+vValue.(map[string]any)["operator"].(string)] = parseValue
			case "regex":
				bsonM["$regex"] = primitive.Regex{
					Pattern: vValue.(map[string]any)["value"].(string),
					Options: vValue.(map[string]any)["regexOption"].(string),
				}
			case "size":
				parseValue, _ := o.scalars["Int"].Set(vValue.(map[string]any)["value"].(string))
				bsonM["$size"] = parseValue
			case "mod":
				divisor := vValue.(map[string]any)["modOption"].(map[string]any)["divisor"].(float64)
				remainder := vValue.(map[string]any)["modOption"].(map[string]any)["remainder"].(float64)
				bsonM["$mod"] = []float64{divisor, remainder}
			default:
				parseValue, _ := o.scalars[vValue.(map[string]any)["kind"].(string)].Set(vValue.(map[string]any)["value"].(string))
				bsonM["$"+vValue.(map[string]any)["operator"].(string)] = parseValue
			}
		}
		bsonE.Value = bsonM
		r = append(r, bsonE)
	}

	if len(and) > 0 {
		bsonE := bson.E{}
		bsonE.Key = "$and"
		bsonA := bson.A{}
		aArgs := args["and"].([]any)
		for _, aArgsValue := range aArgs {
			eResult := o.parse(aArgsValue.(map[string]any))
			bsonA = append(bsonA, eResult)
		}
		bsonE.Value = bsonA
		r = append(r, bsonE)
	}
	if len(not) > 0 {
		bsonE := bson.E{}
		bsonE.Key = "$not"
		bsonA := bson.A{}
		aArgs := args["not"].([]any)
		for _, aArgsValue := range aArgs {
			eResult := o.parse(aArgsValue.(map[string]any))
			bsonA = append(bsonA, eResult)
		}
		bsonE.Value = bsonA
		r = append(r, bsonE)
	}
	if len(nor) > 0 {
		bsonE := bson.E{}
		bsonE.Key = "$nor"
		bsonA := bson.A{}
		aArgs := args["nor"].([]any)
		for _, aArgsValue := range aArgs {
			eResult := o.parse(aArgsValue.(map[string]any))
			bsonA = append(bsonA, eResult)
		}
		bsonE.Value = bsonA
		r = append(r, bsonE)
	}
	if len(or) > 0 {
		bsonE := bson.E{}
		bsonE.Key = "$or"
		bsonA := bson.A{}
		aArgs := args["or"].([]any)
		for _, aArgsValue := range aArgs {
			eResult := o.parse(aArgsValue.(map[string]any))
			bsonA = append(bsonA, eResult)
		}
		bsonE.Value = bsonA
		r = append(r, bsonE)
	}

	bsonE := bson.E{}
	bsonE.Key = args["field"].(string)
	selector := bson.M{}
	cs := bson.A{}
	operator := bson.M{}

	if geoWithinOperators != nil && geoWithinOperators.CenterSphere.Lat != nil {
		centerSphere := geoWithinOperators.CenterSphere
		coordinates := bson.A{*centerSphere.Lat, *centerSphere.Lon}
		cs = append(cs, coordinates)
		cs = append(cs, centerSphere.calc())

		operator["$centerSphere"] = cs
		selector["$geoWithin"] = operator
		bsonE.Value = selector
		r = append(r, bsonE)
	}
	if geoWithinOperators != nil && geoWithinOperators.Box.BottomLeft != nil {
		box := geoWithinOperators.Box
		coordinates1 := bson.A{*&box.BottomLeft.Lat, *&box.BottomLeft.Lon}
		coordinates2 := bson.A{*&box.UpperRight.Lat, *&box.UpperRight.Lon}
		cs = append(cs, coordinates1, coordinates2)
		operator["$box"] = cs
		selector["$geoWithin"] = operator
		bsonE.Value = selector
		r = append(r, bsonE)
	}
	if geoWithinOperators != nil && geoWithinOperators.Center.Lat != nil {
		center := geoWithinOperators.Center
		coordinates := bson.A{center.Lat, center.Lon}
		cs = append(cs, coordinates)
		cs = append(cs, center.deegreeCalc())
		operator["$centerSphere"] = cs
		selector["$geoWithin"] = operator
		bsonE.Value = selector
		r = append(r, bsonE)
	}
	if geoWithinOperators != nil && len(geoWithinOperators.Polygon) > 0 && geoWithinOperators.Polygon[0].Lat != nil {
		polygon := geoWithinOperators.Polygon
		if len(polygon) < 3 {
			r = nil
			return
		}
		for _, coor := range polygon {
			cs = append(cs, coor)
		}
		operator["$polygon"] = cs
		selector["$geoWithin"] = operator
		bsonE.Value = selector
		r = append(r, bsonE)
	}
	if geoSpatialNear != nil && geoSpatialNear.PointCoordinates.Lat != nil {

		cs = append(cs, *geoSpatialNear.PointCoordinates.Lat, *geoSpatialNear.PointCoordinates.Lon)
		selector["$near"] = cs
		if geoSpatialNear.MaxDistance != nil {
			selector["$maxDistance"] = *geoSpatialNear.MaxDistance
		}
		bsonE.Value = selector
		r = append(r, bsonE)

	}
	/*
		if len(x["elemMatch"].([]interface{})) > 0 {
			for _, elemMatch := range x["elemMatch"].([]interface{}) {
				bsonMatch := bson.E{}
				input := elemMatch.(map[string]interface{})
				var value string = input["value"].(string)
				bsonMatch.Key = input["field"].(string)
				bsonMatch.Value = bson.M{"$elemMatch": bson.M{input["fieldMatch"].(string): o.parseValue(input["kind"].(string), &value)}}
				search = append(search, bsonMatch)
			}
		}
	*/
	return
}
