package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mastercard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><circle cx="7" cy="12" r="7"/><circle cx="17" cy="12" r="7" fill-opacity=".8"/></g>`),
		g.Group(children),
	)
}