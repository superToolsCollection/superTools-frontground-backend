package service

import "superTools-frontground-backend/internal/crawler/worker"

/**
* @Author: super
* @Date: 2020-12-31 14:53
* @Description:
**/

type CrawlDataRequest struct {
	Url    string `form:"url" binding:"required,min=2,max=4294967295"`
	Patten string `form:"patten" binding:"required,min=2,max=4294967295"`
}

type ICrawlDataService interface {
	CrawlDataFromURL(param *CrawlDataRequest) ([]string, error)
}

type CrawlService struct {}

func NewCrawlService() ICrawlDataService {
	return &CrawlService{}
}

func (s *CrawlService) CrawlDataFromURL(param *CrawlDataRequest) ([]string, error) {
	r := worker.Request{
		Url:    param.Url,
		Patten: param.Patten,
		Parser: s,
	}
	return worker.Worker(r)
}

func (s *CrawlService) Parse(contents []byte, url string) ([]string, error) {
	
}