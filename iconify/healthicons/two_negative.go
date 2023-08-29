package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthicons2Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM30 18a4 4 0 0 0-4-4h-4a4.002 4.002 0 0 0-3.773 2.666a2 2 0 0 1-3.771-1.332A8.003 8.003 0 0 1 22 10h4a8 8 0 0 1 5.364 13.935L20.948 34H32a2 2 0 1 1 0 4H16a2 2 0 0 1-1.39-3.438l14-13.528l.056-.052A3.985 3.985 0 0 0 30 18Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons2Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}