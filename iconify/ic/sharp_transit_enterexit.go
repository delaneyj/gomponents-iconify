package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTransitEnterexit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 18H6V8h3v4.77L15.98 6L18 8.03L11.15 15H16v3z"/>`),
		g.Group(children),
	)
}