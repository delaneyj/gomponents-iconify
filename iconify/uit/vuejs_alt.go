package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VuejsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.934 3.256a.499.499 0 0 0-.434-.251l-4.99-.003L17.503 3l-4-.026H13.5a.498.498 0 0 0-.43.245l-1.072 1.805l-1.07-1.78A.498.498 0 0 0 10.505 3l-4-.027H6.5A.48.48 0 0 0 6.399 3H1.5a.5.5 0 0 0-.432.752l10.5 18a.5.5 0 0 0 .864 0l10.5-17.995a.5.5 0 0 0 .002-.501zm-12.718.742l1.355 2.259A.5.5 0 0 0 12 6.5h.001a.5.5 0 0 0 .429-.245l1.353-2.28l2.83.02l-3.006 4.917L12 11.54L7.394 3.979l2.822.019zM12 20.508L2.37 4h3.85l5.353 8.76a.493.493 0 0 0 .147.142c.014.01.021.026.035.034a.5.5 0 0 0 .672-.175l5.353-8.759l3.85.002L12 20.508z"/>`),
		g.Group(children),
	)
}