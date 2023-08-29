package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IDocumentsDeniedNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsIDocumentsDeniedNegative0)"><path d="m29 8l6 6h-5a1 1 0 0 1-1-1V8Zm-6.293 12.293a1 1 0 0 0-1.414 1.414L24.586 25l-3.293 3.293a1 1 0 0 0 1.414 1.414L26 26.414l3.293 3.293a1 1 0 0 0 1.414-1.414L27.414 25l3.293-3.293a1 1 0 0 0-1.414-1.414L26 23.586l-3.293-3.293Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm29 5H16a2 2 0 0 0-2 2v30a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V14l-9-9Zm-17 6h-2v27a5 5 0 0 0 5 5h19v-2H15a3 3 0 0 1-3-3V11Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsIDocumentsDeniedNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}