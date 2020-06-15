package emironetwork

import (
	"context"

	"github.com/dominik-robert/emiro/config"
)

type Server struct {
}

func (s *Server) SendQuery(ctx context.Context, query *Query) (*ResponseQueryAnswer, error) {
	answers := searchQueryWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)
	return &ResponseQueryAnswer{QueryAnswers: answers}, nil
}

func (s *Server) SendShow(ctx context.Context, query *Query) (*Answer, error) {
	answer := searchSpecificWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)
	return answer, nil
}

func (s *Server) SendExec(ctx context.Context, query *Query) (*Answer, error) {
	answer := searchSpecificWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)
	return answer, nil
}
