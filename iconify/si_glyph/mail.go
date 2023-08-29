package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.304 3.059H1.701a.65.65 0 0 0-.648.648v8.617a.65.65 0 0 0 .648.648h14.603a.65.65 0 0 0 .649-.648V3.707a.65.65 0 0 0-.649-.648zm-1.398 8.987l-2.884-3.403l-3.009 2.545L5.955 8.57l-2.862 3.477H1.847l3.189-4.353l-3.07-2.6l-.029-1.281l7.076 5.531l7.049-5.62v1.37l-3.017 2.6l3.124 4.338l-1.263.014z"/>`),
		g.Group(children),
	)
}