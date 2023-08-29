package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a3 3 0 0 1 3-3h6.172a3 3 0 0 1 2.12.879l8 8a3 3 0 0 1 0 4.242l-6.17 6.172a3 3 0 0 1-4.243 0l-8-8A3 3 0 0 1 2 11.172V5zm5 1a1 1 0 0 0 0 2h.001a1 1 0 0 0 0-2H7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}