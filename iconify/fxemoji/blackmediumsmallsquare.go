package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackmediumsmallsquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#132028" d="M420 111.249C420 100.662 411.338 92 400.751 92h-291.76c-9.528 0-17.325 7.796-17.325 17.325v293.684c0 9.528 7.796 17.325 17.325 17.325H400.75c10.587 0 19.249-8.662 19.249-19.249V111.249z"/>`),
		g.Group(children),
	)
}