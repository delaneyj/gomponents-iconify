package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComposerEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M25.82 10H30V8h-4.18a3 3 0 0 0-5.64 0H13V5H5v3H2v2h3v3h8v-3h7.18A3 3 0 0 0 22 11.82v7.32A4 4 0 0 0 19.14 22H2v2h17.14a4 4 0 0 0 7.72 0H30v-2h-3.14A4 4 0 0 0 24 19.14v-7.32A3 3 0 0 0 25.82 10zM11 11H7V7h4zm14 12a2 2 0 1 1-2-2a2 2 0 0 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}