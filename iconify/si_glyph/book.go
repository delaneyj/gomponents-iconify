package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 0h2v16H0zm11 6v3h.885L12 6h-1z"/><path d="M3 0v16h10.82c.651 0 1.18-.453 1.18-1.01V1.01C15 .452 14.472 0 13.82 0H3zm10.051 9.053h-.971l-.018 1.01H7.906l.018-1.021h-.967V6.99h.958V5.948l3.042-.01v-.887H7.026v.997h-.997v3.91h1.012v1.017h4.006v1.039H6.961v-.951H5.958v-1.031H4.953V5.991h1.02V4.973h.965V3.938h4.094v1.021h.979v.99h1.041v3.104h-.001z"/><path d="M8 7h2v2H8z"/></g>`),
		g.Group(children),
	)
}