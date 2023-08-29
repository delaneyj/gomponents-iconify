package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceOfWorship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 0l-2 2v2h4V2l-2-2zm-2 4.5L4 6h7L9.5 4.5h-4zM2 6.5a1 1 0 0 0-1 1V13h2V7.5a1 1 0 0 0-1-1zm2 0V13h7V6.5H4zm9 0a1 1 0 0 0-1 1V13h2V7.5a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}