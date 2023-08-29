package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoleFillingCursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="19" cy="19" r="4" fill="currentColor"/><path d="M28 30H10a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h18a2.002 2.002 0 0 1 2 2v18a2.002 2.002 0 0 1-2 2zM10 10v18h18V10z" fill="currentColor"/><path d="M11 2H2v9h2V4h7V2z" fill="currentColor"/>`),
		g.Group(children),
	)
}