package emironetwork

import (
	"context"

	"github.com/dominik-robert/emiro/config"
)

type Server struct {
}

func (s *Server) SendQuery(ctx context.Context, query *Query) (*ResponseQueryAnswer, error) {
	answers, err := searchQueryWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)
	return &ResponseQueryAnswer{QueryAnswers: answers}, err
}

func (s *Server) SendShow(ctx context.Context, query *Query) (*Answer, error) {
	answer, err := searchSpecificWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)
	return answer, err
}

func (s *Server) SendExec(ctx context.Context, query *Query) (*Answer, error) {
	answer, err := searchSpecificWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)
	return answer, err
}

func (s *Server) SendNew(ctx context.Context, query *QueryFull) (*Response, error) {
	answer, err := createNewWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Data)

	return &Response{Succeed: answer}, err
}
