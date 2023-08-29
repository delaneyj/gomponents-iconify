package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsTravelAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm35 9.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0ZM17 18a2 2 0 0 1 2-2h17a6 6 0 0 1 1 11.917V40a2 2 0 1 1-4 0v-9h-3v9a2 2 0 1 1-4 0V20h-7a2 2 0 0 1-2-2Zm20 2a2 2 0 1 1 0 4v-4ZM6 32a2 2 0 0 1 2-2h3v-2a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v2h3a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-8Zm10-2v-2h-3v2h3Zm-6 9v-6h2v6h-2Zm7-6v6h2v-6h-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTravelAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}