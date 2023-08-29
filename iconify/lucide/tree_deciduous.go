package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeDeciduous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19h8a4 4 0 0 0 3.8-2.8a4 4 0 0 0-1.6-4.5c1-1.1 1-2.7.4-4c-.7-1.2-2.2-2-3.6-1.7a3 3 0 0 0-3-3a3 3 0 0 0-3 3c-1.4-.2-2.9.5-3.6 1.7c-.7 1.3-.5 2.9.4 4a4 4 0 0 0-1.6 4.5A4 4 0 0 0 8 19Zm4 0v3"/>`),
		g.Group(children),
	)
}