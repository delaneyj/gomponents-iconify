package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V8a1 1 0 0 0-.293-.707l-4-4A1 1 0 0 0 16 3H4zm6 11c0-1.6 1.333-2 2-2s2 .4 2 2s-1.333 2-2 2s-2-.4-2-2zm4-9H8v3h6V5z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}