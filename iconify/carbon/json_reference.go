package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JsonReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm25-8l-2-6h-2v10h2v-6l2 6h2V6h-2v6zm-7.666-6h-2.667A1.668 1.668 0 0 0 17 7.667v6.667A1.668 1.668 0 0 0 18.666 16h2.667A1.668 1.668 0 0 0 23 14.334V7.667A1.668 1.668 0 0 0 21.334 6zM21 14h-2V8h2zM9 7.667V10a2.002 2.002 0 0 0 2 2h2v2H9v2h4.333A1.668 1.668 0 0 0 15 14.334V12a2.002 2.002 0 0 0-2-2h-2V8h4V6h-4.334A1.668 1.668 0 0 0 9 7.667zM5 14H3v-2H1v2.333A1.668 1.668 0 0 0 2.667 16h2.667A1.668 1.668 0 0 0 7 14.334V6H5z"/>`),
		g.Group(children),
	)
}