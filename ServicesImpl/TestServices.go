package ServicesImpl

import (
	"context"
	"micro/Services"
	"strconv"
)

type TestService struct {
}

func (this *TestService) Call(ctx context.Context, req *Services.TestRequest, resp *Services.TestResponse) error {
	resp.Data = "Test" + strconv.Itoa(int(req.Id))
	return nil
}
