package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0zm0 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13z"/><path fill="currentColor" d="M9 8L5 5v6zm2-3H9v6h2V5z"/>`),
		g.Group(children),
	)
}