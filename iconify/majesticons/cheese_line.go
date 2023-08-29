package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheeseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9v4c-.667 0-2 .4-2 2s1.333 2 2 2v3H4.5C2 15 2 6 11 4l10 5zm0 0H5m3 7h.001M13 13h.001M15 16h.001"/>`),
		g.Group(children),
	)
}