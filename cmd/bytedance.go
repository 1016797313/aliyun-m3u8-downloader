package cmd

import (
	"fmt"
	"log"

	"github.com/ddliu/go-httpclient"
	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/download"
	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/tool"
	"github.com/spf13/cobra"
)

// bytedanceCmd represents the bytedance command
var bytedanceCmd = &cobra.Command{
	Use:   "bytedance",
	Short: "字节跳动，火山引擎视频云视频加密下载工具",
	Long: `字节跳动，火山引擎视频云视频加密下载工具. 使用示例:
aliyun-m3u8-downloader bytedance -p "PlayAuthToken" -o=/data/example --chanSize 1 -f 文件名`,
	PreRun: func(cmd *cobra.Command, args []string) {
		httpclient.Defaults(httpclient.Map{
			"Accept-Encoding":        "gzip",
			httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		})
		if referer, _ := cmd.Flags().GetString("referer"); referer != "" {
			httpclient.Defaults(httpclient.Map{
				httpclient.OPT_REFERER: referer,
			})
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		playAuth, _ := cmd.Flags().GetString("playAuth")
		filename, _ := cmd.Flags().GetString("filename")
		output, _ := cmd.Flags().GetString("output")
		chanSize, _ := cmd.Flags().GetInt("chanSize")
		if playAuth == "" {
			tool.PanicParameter("playAuth")
		}
		if chanSize <= 0 {
			panic("parameter 'chanSize' must be greater than 0")
		}
		if err := download.Bytedance(output, filename, chanSize, playAuth); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(bytedanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aliyunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	bytedanceCmd.Flags().StringP("playAuth", "p", "", "web播放认证信息")
	bytedanceCmd.Flags().StringP("referer", "r", "", "referer请求头")
	bytedanceCmd.Flags().StringP("output", "o", "", "下载保存位置")
	bytedanceCmd.Flags().StringP("filename", "f", "", "保存文件名")
	bytedanceCmd.Flags().IntP("chanSize", "c", 5, "下载并发数")
	_ = aliyunCmd.MarkFlagRequired("playAuth")
}
