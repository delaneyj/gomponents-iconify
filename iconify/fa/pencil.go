package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m363 1408l91-91l-235-235l-91 91v107h128v128h107zm523-928q0-22-22-22q-10 0-17 7l-542 542q-7 7-7 17q0 22 22 22q10 0 17-7l542-542q7-7 7-17zm-54-192l416 416l-832 832H0v-416zm683 96q0 53-37 90l-166 166l-416-416l166-165q36-38 90-38q53 0 91 38l235 234q37 39 37 91z"/>`),
		g.Group(children),
	)
}