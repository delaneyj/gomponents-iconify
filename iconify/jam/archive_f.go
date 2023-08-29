package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 5H1a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-2zm1 2v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7h16zM8 14a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2H8z"/>`),
		g.Group(children),
	)
}