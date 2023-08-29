package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrescriptionDocumentNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPrescriptionDocumentNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM27 4v11a1 1 0 0 0 1 1h11v27a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h17Zm2 .586V14h9.414L29 4.586ZM18 18a1 1 0 0 0-1 1v14h2v-7h1.586l5 5l-3.293 3.293l1.414 1.414L27 32.414l3.293 3.293l1.414-1.414L28.414 31l3.293-3.293l-1.414-1.414L27 29.586l-3.605-3.605A4 4 0 0 0 23 18h-5Zm1 6h4a2 2 0 1 0 0-4h-4v4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPrescriptionDocumentNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}