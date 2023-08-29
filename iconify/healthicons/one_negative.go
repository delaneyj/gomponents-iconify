package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthicons1Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM29 12a2 2 0 0 0-3.414-1.414l-5 5a2 2 0 1 0 2.828 2.828L25 16.828V36a2 2 0 0 0 4 0V12Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons1Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}