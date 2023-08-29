package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 4h-3V0H0v14h6v2h10V7l-3-3zM3 1h4v1H3V1zm12 14H7V5h5v3h3v7zm-2-8V5l2 2h-2z"/>`),
		g.Group(children),
	)
}