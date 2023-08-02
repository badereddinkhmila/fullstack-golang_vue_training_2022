package elastic_searching_products

import (
	"bytes"
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"io"
	"io/ioutil"

	esv7api "github.com/elastic/go-elasticsearch/v7/esapi"

	categoryService "com.fullstack.ecommerce/app/categories/service"
	productModel "com.fullstack.ecommerce/app/products/model"
	productRepo "com.fullstack.ecommerce/app/products/repository"
	userService "com.fullstack.ecommerce/app/users/service"
	"com.fullstack.ecommerce/utils/elastic"
)

type IndexedProduct struct {
	ID                  int     `json:"id"`
	NameProduct         string  `json:"name_product"`
	DescriptionProduct  string  `json:"description_product"`
	Price               float64 `json:"price_product"`
	Quantity            int     `json:"quantity_product"`
	IsAvailable         bool    `json:"is_available"`
	CategoryLable       string  `json:"category_name"`
	CategoryDescription string  `json:"category_description"`
	CreatedBy           string  `json:"provider_name"`
}

func ProductsIndexes() string { return "ProductsIndexes" }

func Index(ctx context.Context, product productModel.Product) error {
	// XXX: Excluding OpenTelemetry and error checking for simplicity
	/*
		ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "Product.Index")
		defer span.End()
	*/

	category, err := categoryService.FindByID(ctx, product.CategoryID)
	if err != nil {
		return err
	}

	provider, err := userService.GetUserById(ctx, product.CreatedBy)
	if err != nil {
		return err
	}

	body := IndexedProduct{
		ID:                  product.ID,
		NameProduct:         product.NameProduct,
		DescriptionProduct:  product.DescriptionProduct,
		Price:               product.Price,
		Quantity:            product.Quantity,
		IsAvailable:         product.IsAvailable,
		CategoryLable:       category.Lable,
		CategoryDescription: category.Description,
		CreatedBy:           provider.Username,
	}

	var buf bytes.Buffer

	_ = json.NewEncoder(&buf).Encode(body) // XXX: error omitted

	req := esv7api.IndexRequest{
		Index:      ProductsIndexes(),
		Body:       &buf,
		DocumentID: string(product.ID),
		Refresh:    "true",
	}

	resp, _ := req.Do(ctx, elastic.Client()) // XXX: error omitted
	defer resp.Body.Close()

	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func Search(ctx context.Context, name *string, description *string, price *[]float64,
	categoryLable *string, categoryDescription *string,
	provider *string) ([]*productModel.Product, error) {

	// XXX: Excluding OpenTelemetry and error checking for simplicity
	if name == nil && description == nil && price == nil &&
		categoryLable == nil && categoryDescription == nil && provider == nil {
		return nil, nil
	}

	should := make([]interface{}, 0, 3)

	if description != nil {
		should = append(should, matchQuery(*description, "description_product"), fuzzyQuery(*description, "description_product"))
	}

	if name != nil {
		should = append(should, matchQuery(*name, "name_product"), fuzzyQuery(*name, "name_product"))
	}

	if categoryLable != nil {
		should = append(should, matchQuery(*categoryLable, "category_name"), fuzzyQuery(*categoryLable, "category_name"))
	}

	if categoryDescription != nil {
		should = append(should, matchQuery(*categoryDescription, "category_description"), fuzzyQuery(*categoryDescription, "category_description"))
	}

	if categoryDescription != nil {
		should = append(should, matchQuery(*categoryDescription, "category_description"), fuzzyQuery(*categoryDescription, "category_description"))
	}

	if provider != nil {
		should = append(should, matchQuery(*provider, "provider_name"), fuzzyQuery(*provider, "provider_name"))
	}

	if price != nil {
		should = append(should, rangeQuery(*price, "WITHIN"))
	}

	var query map[string]interface{}

	if len(should) > 1 {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": should,
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": should[0],
		}
	}

	var buf bytes.Buffer

	_ = json.NewEncoder(&buf).Encode(query)

	req := esv7api.SearchRequest{
		Index: []string{ProductsIndexes()},
		Body:  &buf,
	}

	resp, _ := req.Do(ctx, elastic.Client()) // XXX: error omitted
	defer resp.Body.Close()

	var hits struct {
		Hits struct {
			Hits []struct {
				Source IndexedProduct `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	_ = json.NewDecoder(resp.Body).Decode(&hits) // XXX: error omitted

	res := make([]int, len(hits.Hits.Hits))
	// Getting on the ids of the returned results
	for i, hit := range hits.Hits.Hits {
		res[i] = hit.Source.ID
	}

	result, err := productRepo.FindAllByIds(ctx, res)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateIndex(ctx context.Context, product productModel.Product) error {
	return nil
}

func Delete(ctx context.Context, id string) error {
	// tracing
	/* To do here */

	req := esv7api.DeleteRequest{
		Index:      ProductsIndexes(),
		DocumentID: id,
	}

	resp, err := req.Do(ctx, elastic.Client())
	if err != nil {
		return errors.Wrap(err, "DeleteRequest.Do")
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return errors.Wrapf(errors.NewPlain("Error deleting product index"), "DeleteRequest.Do %v", resp.StatusCode)
	}

	io.Copy(ioutil.Discard, resp.Body)

	return nil
}

func rangeQuery(between []float64, relation string) map[string]interface{} {
	return map[string]interface{}{
		"range": map[string]interface{}{
			"price_product": map[string]interface{}{
				"gte":      between[0],
				"lte":      between[1],
				"relation": relation,
			},
		},
	}
}

func matchQuery(field string, fieldIndex string) map[string]interface{} {
	return map[string]interface{}{
		"match": map[string]interface{}{
			fieldIndex: field,
		},
	}
}

func fuzzyQuery(field string, fieldIndex string) map[string]interface{} {
	return map[string]interface{}{
		"fuzzy": map[string]interface{}{
			fieldIndex: map[string]interface{}{
				"value":     field,
				"fuzziness": "AUTO"},
		},
	}
}
