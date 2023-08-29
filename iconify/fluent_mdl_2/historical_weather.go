package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoricalWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1664h1920v128H0V256h128v441l525-263l340 340l459 115l487-486l90 90l-537 538l-565-141l-300-300l-499 249v297l520-115l519 259h354l434-217l58 114l-462 231h-414l-505-253l-504 112v397z"/>`),
		g.Group(children),
	)
}