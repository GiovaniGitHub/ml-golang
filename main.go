package main

import (
	"os"

	"github.com/GiovaniGitHub/ml-golang/clusters"
	"github.com/GiovaniGitHub/ml-golang/regressions"
)

func main() {
	switch argsWithProg := os.Args[len(os.Args)-1]; argsWithProg {
	case "linear":
		regressions.LinearRegression()

	case "polynomial":
		regressions.PolynomialRegression()

	case "simple":
		regressions.SimpleLinearRegression()

	case "rbf":
		regressions.RBFRegression()

	case "knn":
		clusters.KNNClassifier()
	}

}
