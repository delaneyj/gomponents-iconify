package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zm-7.496-4.868A1 1 0 0 1 17 8v8a1 1 0 0 1-1.496.868L9 13.152V16a1 1 0 1 1-2 0V8a1 1 0 1 1 2 0v2.848l6.504-3.716zM15 9.723L11.016 12L15 14.277V9.723z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}