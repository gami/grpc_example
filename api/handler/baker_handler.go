package handler

import (
	"math/rand"
	"sync"
	"time"
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pancake/maker/gen/api"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//BakerHandler はパンケーキを焼きます
type BakerHandler struct {
	report *report
}

type report struct {
	sync.Mutex // 複数人が同時に焼いても大丈夫にしておきます
	data       map[api.Pancake_Menu]int
}

//NewBakerHandler はBakerHandlerを初期化します
func NewBakerHandler() *BakerHandler {
	return &BakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int),
		},
	}
}

//Bake は指定されたメニューのパンを焼いて、焼けたパンをResponseとして返します。
func (h *BakerHandler) Bake(
	ctx context.Context,
	req *api.BakeRequest,
	) (*api.BakeResponse, error) {
	//バリデーション
	if req.Menu == api.Pancake_UNKNOWN || req.Menu > api.Pancake_SPICY_CURRY{
		return nil, status.Errorf(codes.InvalidArgument, "パンケーキを選んでください！")
	}

	//パンを焼いて、数を記録します
	now := time.Now()
	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	//レスポンスを作って返します
	return &api.BakeResponse{
		Pancake: &api.Pancake{
			Menu:           req.Menu,
			ChefName:       "gami", //ワンオペ
			TechnicalScore: rand.Float32(),
			CreateTime: &timestamp.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}

//Report は、焼けたパンの数を報告します。
func (h *BakerHandler) Report(ctx context.Context, req *api.ReportRequest) (*api.ReportResponse, error) {
	//sliceを初期化
	counts := make([]*api.Report_BakeCount, 0)

	//レポートを作ります
	h.report.Lock()
	for k, v := range h.report.data {
		counts = append(counts, &api.Report_BakeCount{
			Menu:  k,
			Count: int32(v),
		})
	}
	h.report.Unlock()

	//レスポンスを作って返します
	return &api.ReportResponse{
		Report: &api.Report{
			BakeCounts: counts,
		},
	}, nil
}
