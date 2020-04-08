package mysql

import (
	"database/sql"

	parsetmpl "github.com/DoodgeMatvey/rbac-parsing-tmpl"

	_ "github.com/go-sql-driver/mysql"
)

func OpenD(dbInfo string) (*sql.DB, error) {
	return sql.Open("mysql", dbInfo)
}

func InsertData(db *sql.DB, result parsetmpl.Result) {

	stmt, err := db.Prepare("INSERT INTO template (service_name, feature_name, feature_id, feature_descr, endpoints_path, endpoints_methods) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	for _, cFeature := range result.Features {
		for _, endpoints := range cFeature.Endpoints {
			for cPath, method := range endpoints {

				_, err := stmt.Exec(result.ServiceName, cFeature.FeatureName, cFeature.ID, cFeature.Description, cPath, method)
				if err != nil {
					panic(err)
				}

			}
		}
	}

}
