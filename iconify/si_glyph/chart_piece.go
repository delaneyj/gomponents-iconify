package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPiece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.392 1.049c-.349-.051-.562-.033-.925-.033c-4.095 0-7.424 3.334-7.424 7.43s3.329 7.431 7.424 7.431c4.094 0 7.423-3.335 7.423-7.431c0-.346.044-.598-.002-.929l-7.399.961l.903-7.429z"/><path d="m11.375.047l-.945 6.539l6.613-.928C16.638 2.764 14.314.466 11.375.047z"/></g>`),
		g.Group(children),
	)
}