package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12 13H4c-.5 0-1-.5-1-1V5c0-2 .8-4 4-4c3.1 0 4 2 4 4v6h1c.5 0 1 .5 1 1s-.5 1-1 1zm-7-2h4V5c0-1 0-2-2-2S5 4 5 5v6z"/>`),
		g.Group(children),
	)
}