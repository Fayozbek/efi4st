/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"../dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"strconv"
)

func TestResults(ctx iris.Context) {

	testResults := dbprovider.GetDBManager().GetTestResults()

	ctx.ViewData("error", "")

	if len(testResults) < 1 {
		ctx.ViewData("error", "Error: No apps available. Add one!")
	}

	ctx.ViewData("testResultsList", testResults)
	ctx.View("testResults.html")
}

func ShowTestResult(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing firmware Id!")
	}

	testResult := dbprovider.GetDBManager().GetTestResultInfo(i)

	ctx.ViewData("testResult", testResult)
	ctx.View("showTestResult.html")
}


func RemoveTestResult(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveTestResult(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing test Result!")
	}

	testResults := dbprovider.GetDBManager().GetTestResults()

	ctx.ViewData("testResultsList", testResults)
	ctx.View("testResults.html")
}



