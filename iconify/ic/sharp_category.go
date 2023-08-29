package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCategory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2l-5.5 9h11z"/><circle cx="17.5" cy="17.5" r="4.5" fill="currentColor"/><path fill="currentColor" d="M3 13.5h8v8H3z"/>`),
		g.Group(children),
	)
}