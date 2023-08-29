package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsNoNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm-7.778-13.636a1 1 0 0 0 1.414 1.414L24 25.414l6.364 6.364a1 1 0 0 0 1.414-1.414L25.414 24l6.364-6.364a1 1 0 0 0-1.414-1.414L24 22.586l-6.364-6.364a1 1 0 0 0-1.414 1.414L22.586 24l-6.364 6.364Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsNoNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}