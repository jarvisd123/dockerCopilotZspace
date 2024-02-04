package version

import (
	"context"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/utiles"
	"os"
	"time"

	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProgramLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProgramLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProgramLogic {
	return &UpdateProgramLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProgramLogic) UpdateProgram() (resp *types.Resp, err error) {
	resp = &types.Resp{}
	err = utiles.UpdateProgram(l.svcCtx)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		resp.Data = map[string]interface{}{}
		return resp, err
	}
	resp.Code = 200
	resp.Msg = "success"
	go func() {
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}()
	resp.Data = map[string]interface{}{}
	return resp, nil
}
