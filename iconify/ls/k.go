package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func K(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M411 228L137 466l319 288H349L81 513l-9 9v232H0V0h72v426l229-198h110z"/>`),
		g.Group(children),
	)
}