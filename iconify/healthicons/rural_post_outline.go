package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuralPostOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M39 12h4v2h-4v-2Z"/><path fill-rule="evenodd" d="M6 16h17v2H8v19h21V25h8v12h3V18h-1v-2h3v21h1a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2h1V16Zm25 21h4V27h-4v10Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M27 32h-1v-7H11v7h-1v2h17v-2Zm-14-5v5h11v-5H13Z" clip-rule="evenodd"/><path d="M5 12v2h18v-2H5Zm29 3h-2v-2h-2v2h-2v2h2v2h2v-2h2v-2Z"/><path fill-rule="evenodd" d="M37 12a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-8Zm-10 0h8v8h-8v-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}