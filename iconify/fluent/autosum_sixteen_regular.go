package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutosumSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.54 2.304A.5.5 0 0 1 3 2h9.5a.5.5 0 1 1 0 1H4.171l4.074 4.253a.5.5 0 0 1 .024.665L4.063 13H12.5a.5.5 0 1 1 0 1H3a.5.5 0 0 1-.385-.819l4.6-5.558l-4.576-4.777a.5.5 0 0 1-.099-.542Z"/>`),
		g.Group(children),
	)
}