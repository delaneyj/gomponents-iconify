package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLibraryMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H6v16h16V2zm-4 5h-3v5.5a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 2.5-2.5c.57 0 1.08.19 1.5.51V5h4v2zM4 6H2v16h16v-2H4V6z"/>`),
		g.Group(children),
	)
}