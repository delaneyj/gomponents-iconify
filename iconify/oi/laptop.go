package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1.34 1a.5.5 0 0 0-.34.5V5H0v1.5c0 .28.22.5.5.5h7.01c.28 0 .5-.22.5-.5V5h-1V1.5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM2 2h4v3H5v1H3V5H2V2z"/>`),
		g.Group(children),
	)
}