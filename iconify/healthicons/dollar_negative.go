package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DollarNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsDollarNegative0)"><path d="M18 18a4 4 0 0 1 4-4v8a4 4 0 0 1-4-4Zm12 12a4 4 0 0 1-4 4v-8a4 4 0 0 1 4 4Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 6a2 2 0 0 1 2 2v2a8.003 8.003 0 0 1 7.544 5.334a2 2 0 0 1-3.771 1.332A4.002 4.002 0 0 0 26 14v8a8 8 0 1 1 0 16v2a2 2 0 1 1-4 0v-2a8.003 8.003 0 0 1-7.544-5.334a2 2 0 0 1 3.771-1.332A4.002 4.002 0 0 0 22 34v-8a8 8 0 1 1 0-16V8a2 2 0 0 1 2-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDollarNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}