package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 10a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1h-11zm-4-3h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm15 11h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1zm4-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}