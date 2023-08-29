package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMusicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3c-1.1 0-2 .9-2 2v16l7-3l7 3V5a2 2 0 0 0-2-2H7m0 2h10v13l-5-2.18L7 18V5m5 1v5.3c-.3-.2-.6-.3-1-.3a2 2 0 1 0 0 4c1.11 0 2-.89 2-2V8h2V6h-3Z"/>`),
		g.Group(children),
	)
}