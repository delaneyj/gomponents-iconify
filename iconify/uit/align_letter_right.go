package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLetterRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 4h12a.5.5 0 0 0 0-1h-12a.5.5 0 0 0 0 1zm12 17h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm0-4h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1zm0-6h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm0-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}