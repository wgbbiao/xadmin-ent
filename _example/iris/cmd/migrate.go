package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wgbbiao/xadminent"
)

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "迁移数据库",
	Long:  `迁移数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		xadminent.AutoMigrate()
	},
}
