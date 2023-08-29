package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikingDiamonds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M30.875 36.063L88.25 300.625L202.406 400.97l24.688-76.814L30.874 36.062zm102 42.343L333.72 198.344l96.374 2.375l-46.78-68.345l-250.44-53.97zm3.5 42.28l202 284.595l149.5 78.626L440.78 365.78L136.376 120.69z"/>`),
		g.Group(children),
	)
}