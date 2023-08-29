package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5.5A4.5 4.5 0 0 1 6.5 1H13v1h-2v12h-1V2H8v12H7v-4h-.5A4.5 4.5 0 0 1 2 5.5ZM7 9V2h-.5a3.5 3.5 0 1 0 0 7H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}