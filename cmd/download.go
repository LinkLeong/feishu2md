package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/88250/lute"
	"github.com/LinkLeong/feishu2md/core"
	"github.com/LinkLeong/feishu2md/utils"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
	"github.com/pkg/errors"
)

func handleUrlArgument(url string) error {
	configPath, err := core.GetConfigFilePath()
	utils.CheckErr(err)
	config, err := core.ReadConfigFromFile(configPath)
	utils.CheckErr(err)

	reg := regexp.MustCompile("^https://[a-zA-Z0-9-]+.(feishu.cn|larksuite.com)/(docx|wiki)/([a-zA-Z0-9]+)")
	matchResult := reg.FindStringSubmatch(url)
	if matchResult == nil || len(matchResult) != 4 {
		return errors.Errorf("Invalid feishu/larksuite URL format")
	}

	docType := matchResult[2]
	docToken := matchResult[3]
	fmt.Println("Captured document token:", docToken)

	ctx := context.Background()

	client := core.NewClient(
		config.Feishu.AppId, config.Feishu.AppSecret,
	)

	// for a wiki page, we need to renew docType and docToken first
	if docType == "wiki" {
		node, err := client.GetWikiNodeInfo(ctx, docToken)
		utils.CheckErr(err)
		docType = node.ObjType
		docToken = node.ObjToken
	}

	docx, blocks, err := client.GetDocxContent(ctx, docToken)
	utils.CheckErr(err)

	parser := core.NewParser(ctx)

	title := docx.Title
	markdown := parser.ParseDocxContent(docx, blocks)

	for _, imgToken := range parser.ImgTokens {
		localLink, err := client.DownloadImage(ctx, imgToken, config.Output.ImageDir)
		if err != nil {
			return err
		}
		markdown = strings.Replace(markdown, imgToken, localLink, 1)
	}

	engine := lute.New(func(l *lute.Lute) {
		l.RenderOptions.AutoSpace = true
	})
	result := engine.FormatStr("md", markdown)

	mdName := fmt.Sprintf("%s.md", docToken)
	if config.Output.TitleAsFilename {
		mdName = fmt.Sprintf("%s.md", title)
	}
	if err = os.WriteFile(mdName, []byte(result), 0o644); err != nil {
		return err
	}
	fmt.Printf("Downloaded markdown file to %s\n", mdName)

	return nil
}

func getSpace(spaceId string) error {
	configPath, err := core.GetConfigFilePath()
	utils.CheckErr(err)
	config, err := core.ReadConfigFromFile(configPath)
	utils.CheckErr(err)
	client := lark.NewClient(config.Feishu.AppId, config.Feishu.AppSecret)
	req := larkwiki.NewListSpaceNodeReqBuilder().
		SpaceId(spaceId).
		Build()
	// 发起请求
	resp, err := client.Wiki.SpaceNode.List(context.Background(), req)
	// 处理错误
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return err
	}

	// 业务处理
	fmt.Println(resp.Data)
	for _, v := range resp.Data.Items {
		err := download(*v.ObjToken)
		if err != nil {
			fmt.Println(err, "objtoken", *v.ObjToken)
		}
		if *v.HasChild {
			getChildSpace(*v.SpaceId, *v.NodeToken)
		}
	}
	return nil
}
func getChildSpace(spaceId, nodeToken string) error {
	configPath, err := core.GetConfigFilePath()
	utils.CheckErr(err)
	config, err := core.ReadConfigFromFile(configPath)
	utils.CheckErr(err)
	client := lark.NewClient(config.Feishu.AppId, config.Feishu.AppSecret)
	req := larkwiki.NewListSpaceNodeReqBuilder().
		SpaceId(spaceId).
		ParentNodeToken(nodeToken).
		Build()
	// 发起请求
	resp, err := client.Wiki.SpaceNode.List(context.Background(), req)
	// 处理错误
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return err
	}

	// 业务处理
	fmt.Println(resp.Data)
	for _, v := range resp.Data.Items {
		err := download(*v.ObjToken)
		if err != nil {
			fmt.Println(err, "objtoken", *v.ObjToken)
		}
		if *v.HasChild {
			getChildSpace(*v.SpaceId, *v.NodeToken)
		}
	}
	return nil
}
func download(docToken string) error {
	configPath, err := core.GetConfigFilePath()
	utils.CheckErr(err)
	config, err := core.ReadConfigFromFile(configPath)
	client := core.NewClient(
		config.Feishu.AppId, config.Feishu.AppSecret,
	)
	ctx := context.Background()

	docx, blocks, err := client.GetDocxContent(ctx, docToken)
	if err != nil {
		fmt.Println(err)
		return err
	}
	utils.CheckErr(err)

	parser := core.NewParser(ctx)

	title := docx.Title
	markdown := parser.ParseDocxContent(docx, blocks)

	for _, imgToken := range parser.ImgTokens {
		localLink, err := client.DownloadImage(ctx, imgToken, config.Output.ImageDir)
		if err != nil {
			return err
		}
		markdown = strings.Replace(markdown, imgToken, localLink, 1)
	}

	engine := lute.New(func(l *lute.Lute) {
		l.RenderOptions.AutoSpace = true
	})
	result := engine.FormatStr("md", markdown)

	mdName := fmt.Sprintf("%s.md", title+docToken)
	// if config.Output.TitleAsFilename {
	// 	mdName = fmt.Sprintf("%s.md", title)
	// }
	if err = os.WriteFile(mdName, []byte(result), 0o644); err != nil {
		return err
	}
	fmt.Printf("Downloaded markdown file to %s\n", mdName)

	return nil
}
