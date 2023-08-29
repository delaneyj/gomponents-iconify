package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 12a2 2 0 1 0-4 0v24a2 2 0 1 0 4 0V26h12v10a2 2 0 1 0 4 0V12a2 2 0 1 0-4 0v10H18V12Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}