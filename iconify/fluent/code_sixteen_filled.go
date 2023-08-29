package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.905 2.815a.75.75 0 0 1 .38.99l-4 9a.75.75 0 1 1-1.37-.61l4-9a.75.75 0 0 1 .99-.38ZM4.498 5.19a.75.75 0 0 1 .063 1.058L3.003 8l1.558 1.752a.75.75 0 1 1-1.122.996l-2-2.25a.75.75 0 0 1 0-.996l2-2.25A.75.75 0 0 1 4.5 5.19Zm7.004 0a.75.75 0 0 1 1.059.062l2 2.25a.75.75 0 0 1 0 .996l-2 2.25a.75.75 0 0 1-1.122-.996L12.996 8L11.44 6.248a.75.75 0 0 1 .063-1.058Z"/>`),
		g.Group(children),
	)
}