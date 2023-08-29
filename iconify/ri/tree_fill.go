package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 7c0 .262-.017.52-.05.774A6 6 0 0 1 13 18.658V22h-2v-3.6a5.5 5.5 0 0 1-6.517-8.657A8.006 8.006 0 0 0 6.286 12.6l1.428-1.4A5.989 5.989 0 0 1 6 7a6 6 0 1 1 12 0Z"/>`),
		g.Group(children),
	)
}