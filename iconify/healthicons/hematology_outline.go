package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HematologyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M39 8H9a1 1 0 0 0-1 1v30a1 1 0 0 0 1 1h30a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1ZM9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Z" clip-rule="evenodd"/><path d="M32.012 27.13c.003 4.4-3.514 7.86-7.994 7.863c-4.48.004-8.003-3.45-8.006-7.85c-.004-4.4 7.988-14.15 7.988-14.15s8.008 10.051 8.012 14.137Z"/></g>`),
		g.Group(children),
	)
}