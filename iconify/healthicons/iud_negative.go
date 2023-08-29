package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IudNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsIudNegative0)"><path d="M22 40a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM11 6a1 1 0 1 1 0-2h12a1 1 0 0 1 1 1a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H25v30.126A4.002 4.002 0 0 1 24 44a4 4 0 0 1-1-7.874V6H11Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsIudNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}