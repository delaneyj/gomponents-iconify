package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloodLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 17.472A5.978 5.978 0 0 0 20 19h2v2h-2a7.964 7.964 0 0 1-4-1.07A7.96 7.96 0 0 1 12 21a7.964 7.964 0 0 1-4-1.07A7.96 7.96 0 0 1 4 21H2v-2h2c1.537 0 2.94-.578 4-1.528A5.978 5.978 0 0 0 12 19c1.537 0 2.94-.578 4-1.528Zm-3.427-15.94l.1.08L23 11h-3v6c-.701 0-1.374-.12-2-.341V9.157l-6-5.455l-6 5.454l.001 7.502a5.98 5.98 0 0 1-1.702.335L4 17v-6H1l10.327-9.388a1 1 0 0 1 1.246-.08Z"/>`),
		g.Group(children),
	)
}