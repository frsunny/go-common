package feed

import (
	"context"
	"strconv"

	"go-common/app/interface/main/app-card/model"
	"go-common/app/interface/main/app-card/model/card/operate"
)

func (s *Service) SearchFollow(c context.Context, platform, mobiApp, device, buvid string, build int, mid int64) (follow *operate.Follow, err error) {
	const _title = "人气UP主推荐"
	ups, trackID, err := s.search.Follow(c, platform, mobiApp, device, buvid, build, mid)
	if err != nil {
		return
	}
	items := make([]*operate.Follow, 0, len(ups))
	for _, up := range ups {
		if up.Mid != 0 {
			item := &operate.Follow{Pid: up.Mid, Goto: model.GotoMid}
			items = append(items, item)
		}
	}
	if len(items) < 3 {
		return
	}
	id, _ := strconv.ParseInt(trackID, 10, 64)
	if id < 1 {
		return
	}
	follow = &operate.Follow{ID: id, Items: items, Title: _title, Type: "upper"}
	return
}

func (s *Service) SearchFollow2(c context.Context, platform, mobiApp, device, buvid string, build int, mid int64) (follow *operate.Card, err error) {
	const _title = "人气UP主推荐"
	ups, trackID, err := s.search.Follow(c, platform, mobiApp, device, buvid, build, mid)
	if err != nil {
		return
	}
	items := make([]*operate.Card, 0, len(ups))
	for _, up := range ups {
		if up.Mid != 0 {
			item := &operate.Card{ID: up.Mid, Goto: model.GotoMid, Param: strconv.FormatInt(up.Mid, 10), URI: strconv.FormatInt(up.Mid, 10), Desc: up.RecReason}
			items = append(items, item)
		}
	}
	if len(items) < 3 {
		return
	}
	id, _ := strconv.ParseInt(trackID, 10, 64)
	if id < 1 {
		return
	}
	follow = &operate.Card{ID: id, Param: trackID, Items: items, Title: _title, CardGoto: model.CardGotoSearchSubscribe}
	return
}
