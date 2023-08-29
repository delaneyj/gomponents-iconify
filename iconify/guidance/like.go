package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Like(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 9v12m17-.5v-.16a16 16 0 0 1 3.761-10.307l.239-.283V9.5h-9V6A3.5 3.5 0 0 0 10 2.5h-.5V5A4.5 4.5 0 0 1 5 9.5h-.5v11h14Z"/>`),
		g.Group(children),
	)
}