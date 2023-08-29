package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.917 6C6.857 6 6 6.863 6 7.93c0 1.065.857 1.928 1.917 1.928A1.921 1.921 0 0 0 9.832 7.93C9.832 6.864 8.974 6 7.917 6z"/><path d="m10.979 15.11l3.531-4.607l.69.902a7.979 7.979 0 0 0 .799-3.459c0-4.447-3.582-7.946-8-7.946s-8 3.499-8 7.946c0 4.448 3.582 8.054 8 8.054c1.367 0 2.643-.3 3.766-.89h-.786zm1.444-12.419l.846.846l-1.788 1.787l-.844-.845l1.786-1.788zM3.548 13.258l-.845-.846l1.787-1.787l.846.845l-1.788 1.788zM8 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/><path d="m16 13.966l-1.489-1.86l-1.49 1.86h.997v2.003h.97v-2.003H16z"/></g>`),
		g.Group(children),
	)
}