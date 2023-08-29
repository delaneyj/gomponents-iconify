package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBookmarkAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7h-2v2h-2V7h-2V5h2V3h2v2h2v2zm-2 14l-7-3l-7 3V3h9a5.002 5.002 0 0 0 5 7.9V21z"/>`),
		g.Group(children),
	)
}