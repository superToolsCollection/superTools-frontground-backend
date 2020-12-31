package worker

import "superTools-frontground-backend/internal/crawler/fetcher"

/**
* @Author: super
* @Date: 2020-08-16 07:55
* @Description:
**/
func Worker(r Request) ([]string, error){
	contents, _ := fetcher.Fetch(r.Url)
	return r.Parser.Parse(contents, r.Url)
}
