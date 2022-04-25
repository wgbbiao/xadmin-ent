package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wgbbiao/xadminent"
)

var create_super_user = &cobra.Command{
	Use:   "createSuperUser",
	Short: "创建超级管理员",
	Long:  `创建超级管理员`,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := xadminent.CreateAdmin()
		if err != nil {
			panic(err)
		}
		println("创建超级管理员成功", u.ID)
	},
}
