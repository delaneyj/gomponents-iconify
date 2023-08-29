package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.854 8H6.116C4.95 8 5.001 7 5.001 6.903V1c0-1 .083-1 1.115-1h5.738c1.166 0 1.115 0 1.115 1v5.903C12.969 8.06 13 8 11.854 8zM6 1v6.031h6V1H6z"/><path d="M8 8h1.917v6.281H8zm1.969-5v-.969H8V6h1V4.984h.969v-.968H9V3h.969zM10 3h.938v.969H10z"/><path d="M14 15.959H4c0-1.071 2.238-1.938 5-1.938c2.761 0 5 .867 5 1.938z"/></g>`),
		g.Group(children),
	)
}