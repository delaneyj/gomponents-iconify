package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinistryOfHealthNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsMinistryOfHealthNegative0)"><path d="M25 12v-2h-2v2h-2v2h2v2h2v-2h2v-2h-2Zm1 17a2 2 0 1 0-4 0v8h4v-8Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.625 5.22a1 1 0 0 0-1.25 0L8.65 17H6v6h36v-6h-2.65L24.626 5.22ZM20 29a4 4 0 0 1 8 0v8h6v-4h2v-8h4v8h2v4h2v6H4v-6h2v-4h2v-8h4v8h2v4h6v-8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsMinistryOfHealthNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}