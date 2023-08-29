package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BenzeneRingLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M9.98 4.88C10.965 4.292 11.458 4 12 4c.541 0 1.034.293 2.02.88l2.96 1.758c.985.586 1.479.88 1.75 1.362c.27.483.27 1.069.27 2.241v3.518c0 1.172 0 1.758-.27 2.241c-.271.483-.765.776-1.75 1.362l-2.96 1.759c-.986.586-1.479.879-2.02.879c-.541 0-1.034-.293-2.02-.88l-2.96-1.758c-.985-.586-1.479-.88-1.75-1.362C5 15.517 5 14.931 5 13.759V10.24C5 9.07 5 8.483 5.27 8c.271-.483.765-.776 1.75-1.362l2.96-1.759Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 8L2 6m17 2l3-2M5 16l-3 2" opacity=".5"/><path stroke-linecap="round" d="m12 16.884l4-2.384" opacity=".5"/></g>`),
		g.Group(children),
	)
}