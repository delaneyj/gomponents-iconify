package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 7h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1zm17 4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm-2 5h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}