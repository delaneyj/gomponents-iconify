package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.17 5h5.66a3.001 3.001 0 0 0-5.66 0ZM7 6a5 5 0 0 1 10 0v1H7V6ZM5 4v2.586L6.414 8h11.172L19 6.586V4h2v3.414l-3 3V12h3v2h-3v1.586l3 3V22h-2v-2.586l-1.36-1.36A6.009 6.009 0 0 1 13 21.917V22h-1a6.002 6.002 0 0 1-5.64-3.946L5 19.414V22H3v-3.414l3-3V14H3v-2h3v-1.586l-3-3V4h2Zm3 6.414V16a4.002 4.002 0 0 0 3 3.874V10H8.414L8 10.414ZM13 10v9.874c1.725-.444 3-2.01 3-3.874v-5.586L15.586 10H13Z"/>`),
		g.Group(children),
	)
}