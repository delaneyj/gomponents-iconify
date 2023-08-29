package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeSantizerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 6a2 2 0 0 0-2 2v7.1A5.002 5.002 0 0 0 8 20v17a5 5 0 0 0 5 5h12a5.001 5.001 0 0 0 4.584-3h-2.348c-.55.614-1.348 1-2.236 1H13a3 3 0 0 1-3-3v-1h10a1 1 0 1 0 0-2H10v-9h8a1 1 0 1 0 0-2h-8v-3a3 3 0 0 1 3-3h12a2.99 2.99 0 0 1 2.236 1h2.348A5.01 5.01 0 0 0 26 15.1V8a2 2 0 0 0-2-2H14Zm10 2H14v7h10V8Z"/><path d="M37.34 33.69A12.35 12.35 0 0 0 40 25.997V22.4a2.41 2.41 0 0 0-.696-1.697A2.362 2.362 0 0 0 37.627 20H23.374c-.63 0-1.233.253-1.678.703A2.413 2.413 0 0 0 21 22.4v3.598a12.35 12.35 0 0 0 2.66 7.69A12.122 12.122 0 0 0 30.5 38a12.122 12.122 0 0 0 6.84-4.31ZM32 25a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2Z"/></g>`),
		g.Group(children),
	)
}