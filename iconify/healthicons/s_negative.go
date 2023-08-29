package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsSNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 18a4 4 0 0 1 4-4h4a4.002 4.002 0 0 1 3.773 2.666a2 2 0 0 0 3.771-1.332A8.003 8.003 0 0 0 26 10h-4a8 8 0 1 0 0 16h4a4 4 0 0 1 0 8h-4a4.002 4.002 0 0 1-3.773-2.666a2 2 0 0 0-3.771 1.332A8.003 8.003 0 0 0 22 38h4a8 8 0 1 0 0-16h-4a4 4 0 0 1-4-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsSNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}