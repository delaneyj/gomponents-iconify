package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsTwoVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.001 8.2a2.2 2.2 0 1 0-.002-4.4a2.2 2.2 0 0 0 .002 4.4zm0 3.6a2.2 2.2 0 1 0 0 4.402a2.2 2.2 0 0 0 0-4.402z"/>`),
		g.Group(children),
	)
}