package downloader

import (
	"errors"
	"net/http"

	"github.com/maskwang/pholcus/app/downloader/request"
	"github.com/maskwang/pholcus/app/downloader/surfer"
	"github.com/maskwang/pholcus/app/spider"
	"github.com/maskwang/pholcus/config"
	httpdo "github.com/546669204/golang-http-do"
)

type Surfer struct {
	surf    surfer.Surfer
	phantom surfer.Surfer
}

var (
	cookieJar    = &httpdo.Jar{}
	SurferDownloader = &Surfer{
		surf:    surfer.New(cookieJar),
		phantom: surfer.NewPhantom(config.PHANTOMJS, config.PHANTOMJS_TEMP, cookieJar),
	}
)

func (self *Surfer) Download(sp *spider.Spider, cReq *request.Request) *spider.Context {
	ctx := spider.GetContext(sp, cReq)

	var resp *http.Response
	var err error

	switch cReq.GetDownloaderID() {
	case request.SURF_ID:
		resp, err = self.surf.Download(cReq)

	case request.PHANTOM_ID:
		resp, err = self.phantom.Download(cReq)
	}

	if resp.StatusCode >= 400 {
		err = errors.New("响应状态 " + resp.Status)
	}

	ctx.SetResponse(resp).SetError(err)

	return ctx
}
