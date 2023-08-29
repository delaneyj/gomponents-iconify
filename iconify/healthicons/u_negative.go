package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsUNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 12a2 2 0 1 0-4 0v16c0 5.523 4.477 10 10 10s10-4.477 10-10V12a2 2 0 1 0-4 0v16a6 6 0 0 1-12 0V12Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}