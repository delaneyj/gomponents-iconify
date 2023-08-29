package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HematologyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHematologyNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm26.012 18.13c.003 4.4-3.514 7.86-7.994 7.863c-4.48.004-8.003-3.45-8.006-7.85c-.004-4.4 7.988-14.15 7.988-14.15s8.008 10.051 8.012 14.137Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHematologyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}