package cmd

import (
	"github.com/jianggushi/topstory/pkg/spiders"
	_ "github.com/jianggushi/topstory/pkg/spiders/weibo"
	_ "github.com/jianggushi/topstory/pkg/spiders/zhihu"
	"github.com/urfave/cli"
)

var Spider = cli.Command{
	Name:        "spider",
	Usage:       "start spider server",
	Description: "",
	Action:      runSpider,
	Flags:       []cli.Flag{},
}

func runSpider(c *cli.Context) error {
	spiders.RunSpider()
	return nil
}
