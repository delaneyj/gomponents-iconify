package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationArrowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 113.58a15.76 15.76 0 0 1-11.29 15l-76.56 23.56l-23.56 76.56a15.77 15.77 0 0 1-15 11.29h-.3a15.77 15.77 0 0 1-15.07-10.67L33 53.41a1 1 0 0 1-.05-.16a16 16 0 0 1 20.3-20.35l.16.05l175.92 65.26A15.78 15.78 0 0 1 240 113.58Z"/>`),
		g.Group(children),
	)
}