package elastic

import (
	"context"
	"fmt"
	"go-backend/internal/common/env"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9"
)

type Elastic struct {
	EsClient          *elasticsearch.TypedClient
	articleRepository repository.ArticleRepository
	userRepository    repository.UserRepository
}

func NewElastic(env *env.Env, articleRepository repository.ArticleRepository, userRepository repository.UserRepository) *Elastic {
	esClient, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(env.ElasticAddrs),
		elasticsearch.WithBasicAuth(env.ElasticUser, env.ElasticPassword),
		elasticsearch.WithCertificateFingerprint(env.ElasticCertFingerprint),
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	info, err := esClient.Info().Do(context.Background())
	if err != nil {
		log.Fatal("❌ connect elastic error", err)
		return nil
	}

	fmt.Println("✅ [ELASTIC] Connection To Elastic Successfully", info.Name)

	return &Elastic{
		EsClient:          esClient,
		articleRepository: articleRepository,
		userRepository:    userRepository,
	}
}

func (e *Elastic) InitArticle() {
	ctx := context.Background()
	articles, err := e.articleRepository.GetAll(
		ctx,
		pagination.Query{
			Page:     1,
			PageSize: 9999999,
		},
		dto.ArticleFindAllFilters{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, article := range articles {
		fmt.Println("article", article)
		_, err := e.EsClient.
			Index("articles").
			Id(strconv.Itoa(article.ID)).
			Document(article).
			Do(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func (e *Elastic) InitUser() {
	ctx := context.Background()
	users, err := e.userRepository.GetAll(
		ctx,
		pagination.Query{
			Page:     1,
			PageSize: 9999999,
		},
		dto.UserFindAllFilters{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println("user", user)
		_, err := e.EsClient.
			Index("users").
			Id(strconv.Itoa(user.ID)).
			Document(user).
			Do(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
