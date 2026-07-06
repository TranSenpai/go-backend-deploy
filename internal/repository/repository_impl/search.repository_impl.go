package repository_impl

import (
	"context"
	"fmt"
	"go-backend/internal/common/elastic"
	"go-backend/internal/repository"

	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type searchRepository struct {
	elastic *elastic.Elastic
}

func NewSearchRepository(elastic *elastic.Elastic) repository.SearchRepository {
	return &searchRepository{
		elastic: elastic,
	}
}

// FindAll implements [repository.SearchRepository].
func (a *searchRepository) FindAll(ctx context.Context, textSearch string) (any, error) {
	fmt.Println("textSearch", textSearch)
	result, err := a.elastic.EsClient.
		Search().
		Index("articles,users").
		Request(
			&search.Request{
				Query: &types.Query{
					MultiMatch: &types.MultiMatchQuery{
						Query: textSearch,
						Fields: []string{
							"title",
							"content",
							"email",
							"fullName",
						},
						// Nghĩa là với nhiều từ khoá, chỉ cần khớp một phần cũng được
						Operator: &operator.Or,

						// cho phép user gõ sai nhẹ vẫn tìm
						Fuzziness: "AUTO",

						// document nên khớp khoảng 60% số từ
						MinimumShouldMatch: "60%",
					},
				},
			},
		).
		Do(ctx)
	return result, err
}
