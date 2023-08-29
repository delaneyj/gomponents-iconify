package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperShredder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.969 3.003H7.985V0H3v6.982h9.969V3.003z"/><path d="M9.111 0v1.938h3.858L9.111 0zM3 13v2.986h.978V13H3zm2.013 0v3h.974v-3h-.974zM7 13v2.998h.998V13H7zm2 0v2.993h.986V13H9zm2 0v2.964h.963V13H11zM0 8.031V12h15.958V8.031H0zm2.969 2h-1V9h1v1.031zm1 0V9h1.062v1.031H3.969z"/></g>`),
		g.Group(children),
	)
}