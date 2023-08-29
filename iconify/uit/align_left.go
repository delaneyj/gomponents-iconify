package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 7h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm0 4h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1zm15 7h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1zm4-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}