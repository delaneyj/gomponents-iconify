package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntryNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsEntryNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm21 10h-7v27h7V10Zm0-2h-9v31h9v3.613a1 1 0 0 0 1.316.948l13-4.333a1 1 0 0 0 .684-.949V9.721a1 1 0 0 0-.684-.949l-13-4.333A1 1 0 0 0 21 5.387V8Zm6 15c0 1.105-.448 2-1 2s-1-.895-1-2s.448-2 1-2s1 .895 1 2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsEntryNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}