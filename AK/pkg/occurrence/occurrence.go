package occurrence

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"

	pb "google.golang.org/protobuf/proto"

	"github.com/Doozers/adapterKitty/AK/proto"
)

func WikiRequest(ctx context.Context, req *proto.AdapterRequest) (*proto.AdapterResponse, error) {
	formattedRequest := &proto.WikiRequest{}

	if err := pb.Unmarshal(req.Payload, formattedRequest); err != nil {
		return nil, err
	}
	resp, err := http.Get("https://en.wikipedia.org/wiki/" + formattedRequest.GetKeyword())
	if err != nil {
		return nil, err
	}
	fmt.Println(formattedRequest)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r, err := regexp.Compile(formattedRequest.GetNeedle())
	if err != nil {
		return nil, err
	}

	p := &proto.WikiResponse{
		Error:      false,
		Occurrence: int32(len(r.FindAllString(string(body), -1))),
	}

	b, err := pb.Marshal(p)
	if err != nil {
		return nil, err
	}

	fmt.Println("return ->", p)
	return &proto.AdapterResponse{Payload: b}, nil
}
