package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12 5h-2V3H2v4a4 4 0 0 0 7.45 2H12a2 2 0 1 0 0-4zm0 3H9.86A4 4 0 0 0 10 7V6h2a1 1 0 1 1 0 2zm-2 4.5a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1 0-1h7a.5.5 0 0 1 .5.5z"/>`),
		g.Group(children),
	)
}