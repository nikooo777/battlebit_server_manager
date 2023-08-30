package cmd

import (
	"database/sql"
	"fmt"
	"sync"

	"battlebit_admin_panel/config"
	"battlebit_admin_panel/db"
	http "battlebit_admin_panel/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Addr string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&Addr, "addr", ":6456", "address to bind and listen on")

	//Bind to Viper
	err := viper.BindPFlags(serveCmd.PersistentFlags())
	if err != nil {
		logrus.Panic(err)
	}
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs battlebit manager server",
	Long:  `Runs battlebit manager server`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitializeConfiguration()
		stopGrp := &sync.WaitGroup{}
		battlebitDsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", viper.GetString("battlebit_db.user"), viper.GetString("battlebit_db.password"), viper.GetString("battlebit_db.host"), viper.GetString("battlebit_db.database"))
		err := db.Init(battlebitDsn, true)
		if err != nil {
			logrus.Fatalf("failed to initialize database: %s", err.Error())
		}
		db, err := sql.Open("mysql", battlebitDsn)
		if err != nil {
			panic(err.Error())
		}
		sbDSN := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", viper.GetString("sb_db.user"), viper.GetString("sb_db.password"), viper.GetString("sb_db.host"), viper.GetString("sb_db.database"))
		sbDb, err := sql.Open("mysql", sbDSN)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		s := http.NewServer(stopGrp, db, sbDb)
		err = s.Start(Addr)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}
