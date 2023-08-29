package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M0 4.5v20h8.5V27h3v-2.5H20v-20zm5 5h10v10H5ZM22.5 13v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm3 6v6h3v-6zm0 9v6h3v-6zm-80 2v6h3v-6zm80 7v6h3v-6zm-80 2v6h3v-6zm80 7v6h3v-6zm-80 2v6h3v-6zm80 7v6h3v-6zm-80 2v6h3v-6zm80 7v6h3v-6zm-80 2v6h3v-6zm80 7v2.5H80v20h20v-20h-8.5V73Zm-80 2v6h3v-6ZM85 80.5h10v10H85ZM8.5 84v3h6v-3H10zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3zm9 0v3h6v-3z" color="currentColor"/>`),
		g.Group(children),
	)
}