package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 0v32h32V0zm25.599 12.803h-6.401v6.395h-6.395v6.401H6.402V6.402h19.197z"/>`),
		g.Group(children),
	)
}