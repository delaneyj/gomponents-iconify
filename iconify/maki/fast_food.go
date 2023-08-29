package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 8a1 1 0 0 1-1 1H2a1 1 0 1 1 0-2h11a1 1 0 0 1 1 1zM3.5 10H2a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3H3.5zM3 6H2V4a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v2H3zm8-1.5a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0zm-2-1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0zm-2 1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0zm-2-1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0zm-2 1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0z"/>`),
		g.Group(children),
	)
}