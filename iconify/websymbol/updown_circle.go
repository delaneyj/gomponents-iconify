package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpdownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 500q0 136-67 251T751 933t-251 67t-251-67T67 751T0 500t67-251T249 67T500 0t251 67t182 182t67 251zM647 794l162-236H706V265H588v293H485zM412 441h103L353 176L191 441h103v294h118V441z"/>`),
		g.Group(children),
	)
}