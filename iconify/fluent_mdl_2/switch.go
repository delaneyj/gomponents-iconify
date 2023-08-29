package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1408v128H250l163 163l-90 90L6 1472l317-317l90 90l-163 163h1798zm-413-605l163-163H0V512h1798l-163-163l90-90l317 317l-317 317l-90-90z"/>`),
		g.Group(children),
	)
}