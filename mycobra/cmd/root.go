package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short desc",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		//打印flag
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.PersistentFlags().Lookup("config").Value,
			cmd.PersistentFlags().Lookup("license").Value,
			cmd.Flags().Lookup("source").Value,
		)

		fmt.Println("-------------------")
		fmt.Println(
			viper.GetString("author"),
			viper.GetString("license"),
		)
		fmt.Println("root cmd run end")
	},
	TraverseChildren: true,
}

func Execute() {
	rootCmd.Execute()
}

var cfgFile string
var userLicense string

func init() {
	//持久化标志，当前命令及当前命令的所有下级命令都可以使用

	cobra.OnInitialize(initConfig)
	//按名称接收命令行参数
	rootCmd.PersistentFlags().Bool("viper", true, "")
	//指定flag缩写
	rootCmd.PersistentFlags().StringP("author", "a", "YOU NAME", "")
	//通过指针，将值赋值到字段
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	//通过指针，将值赋值到字段，并指定flag缩写
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")

	//本地标志，只能当前命令使用
	//添加本地标志
	rootCmd.Flags().StringP("source", "s", "", "")

	//配置绑定
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	viper.SetDefault("author", "default author")
	viper.SetDefault("license", "default license")

}

// 配置文件初始化
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	fmt.Println("using config file:", viper.ConfigFileUsed())

}
