package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 2.089V24l6.545-6.26L13.089 24V2.089A2.11 2.11 0 0 0 10.98 0l-.077.001h.004h-8.724L2.11 0A2.109 2.109 0 0 0 .001 2.088v.001z"/>`),
		g.Group(children),
	)
}