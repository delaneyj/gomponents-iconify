package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBatterySixty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M17 4h-3V2h-4v2H7v7h10V4z"/><path fill="currentColor" d="M7 11v11h10V11H7z"/>`),
		g.Group(children),
	)
}