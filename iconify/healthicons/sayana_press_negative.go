package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SayanaPressNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsSayanaPressNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm23 6l2-2v7h-2V6Zm-1 13h1v-8h-2a1 1 0 0 0-1 1v5h-1v2h-3a4 4 0 0 0-4 4v17a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V23a4 4 0 0 0-4-4h-3v-2h-1v-5a1 1 0 0 0-1-1h-2v8h1v1a1 1 0 0 1-1 1v3h-2v-3a1 1 0 0 1-1-1v-1Zm2 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm-1 2v2h2v-2h-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsSayanaPressNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}