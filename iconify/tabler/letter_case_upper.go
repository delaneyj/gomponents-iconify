package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterCaseUpper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19V8.5a3.5 3.5 0 0 1 7 0V19m-7-6h7m4 6V8.5a3.5 3.5 0 0 1 7 0V19m-7-6h7"/>`),
		g.Group(children),
	)
}