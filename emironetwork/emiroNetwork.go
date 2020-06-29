package emironetwork

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os/exec"
	"strings"

	"github.com/dominik-robert/emiro/config"
	"google.golang.org/grpc"
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

func (s *Server) SendNew(ctx context.Context, query *QueryFull) (*Response, error) {
	answer, err := createNewWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Data)
	return &Response{Succeed: answer}, err
}

func (s *Server) SendDelete(ctx context.Context, query *Query) (*Response, error) {
	answer, err := deleteWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query)
	return &Response{Succeed: answer}, err
}

func (s *Server) SendExec(ctx context.Context, query *Query) (*Answer, error) {
	answer, err := searchSpecificWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)

	if answer == nil {
		return nil, errors.New("Found no entry in the database for given Query: " + query.Query)
	}

	for _, value := range query.Parameter {
		key, value := splitKeyValue(value)
		answer.Command = strings.ReplaceAll(answer.Command, key, value)
	}
	for key, value := range answer.Params {
		answer.Command = strings.ReplaceAll(answer.Command, key, value)
	}

	answer.Command = query.Prepend + " " + answer.Command + " " + query.Append

	return answer, err
}

func splitKeyValue(keyVal string) (key, value string) {
	arr := strings.Split(keyVal, "=")

	return arr[0], arr[1]
}

func (s *Server) SendExecToRemote(answer *Answer, stream Emiro_SendExecToRemoteServer) error {
	cmd := exec.Command(answer.Path, "-c", answer.Command)

	stdout, errOut := cmd.StdoutPipe()
	stdErr, errErr := cmd.StderrPipe()

	if errOut != nil {
		log.Fatalf("Cannot connect to commands stdOut: %s", errOut)
		return errOut
	}
	if errErr != nil {
		log.Fatalf("Cannot connect to commands stdErr: %s", errErr)
		return errErr
	}

	cmd.Start()

	go func() {
		scannerErr := bufio.NewScanner(stdErr)
		for scannerErr.Scan() {
			m := scannerErr.Text()
			resp := &Response{
				Succeed: true,
				Data:    []byte(m),
			}
			stream.Send(resp)
		}
	}()

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
		m := scannerOut.Text()
		resp := &Response{
			Succeed: true,
			Data:    []byte(m),
		}
		stream.Send(resp)
	}
	cmd.Wait()

	stream.Send(&Response{Succeed: false, Data: []byte("\n")})

	return nil
}

func (s *Server) ExecRemote(query *Query, stream Emiro_ExecRemoteServer) error {
	answer, err := searchSpecificWithElastic(config.Config.GetString("databaseHost"), config.Config.GetInt("databasePort"), config.Config.GetString("databaseIndex"), config.Config.GetBool("databaseInsecure"), query.Query, query.Count, query.All)

	if answer == nil {
		return errors.New("Found no entry in the database for given Query: " + query.Query)
	}

	for _, value := range query.Parameter {
		key, value := splitKeyValue(value)
		answer.Command = strings.ReplaceAll(answer.Command, key, value)
	}
	for key, value := range answer.Params {
		answer.Command = strings.ReplaceAll(answer.Command, key, value)
	}

	answer.Command = query.Prepend + " " + answer.Command + " " + query.Append

	var conn *grpc.ClientConn

	conn, err = grpc.Dial(query.RemoteHost, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to server: %s", err)
	}

	c := NewEmiroClient(conn)

	response, err := c.SendExecToRemote(context.Background(), answer)

	if err != nil {
		log.Println(err)
		return err
	}

	responseTmp, err := response.Recv()
	if err != nil {
		return err
	}

	for responseTmp.Succeed == true {
		stream.Send(responseTmp)
		responseTmp, err = response.Recv()

		if err != nil {
			log.Fatalf("Error while receiving stream: %s", err)
		}
	}

	return nil
}
