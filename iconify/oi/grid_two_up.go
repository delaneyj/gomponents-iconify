package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridTwoUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v3h3V0H0zm5 0v3h3V0H5zM0 5v3h3V5H0zm5 0v3h3V5H5z"/>`),
		g.Group(children),
	)
}