package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.02.002l3.49 7.988l7.987 3.49l-7.988 3.49l-3.49 7.989l-3.49-7.989l-7.988-3.49l7.988-3.49L12.02.002Zm9.304 3.323l-1.568.78l1.568.781l.781 1.57l.781-1.57l1.57-.78l-1.57-.781l-.78-1.57l-.782 1.57ZM12.02 4.998l-1.97 4.511l-4.512 1.971l4.511 1.971l1.971 4.512l1.971-4.512l4.512-1.97l-4.512-1.972l-1.97-4.511Zm7.316 9.758l1.3 2.61l2.61 1.3l-2.61 1.299l-1.3 2.61l-1.3-2.61l-2.61-1.3l2.61-1.299l1.3-2.61Z"/>`),
		g.Group(children),
	)
}