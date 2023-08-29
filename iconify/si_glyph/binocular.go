package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binocular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 14h6v.916H0zm10 0h6v.916h-6zM2.041 3.012L.666 12.917h4.667l-.687-3.905h1.323v-6h-.985V1.979h.981v-.95H2.006v.95H3v1.033h-.959zm8.959 0h-.969v6h1.344l-.709 3.905h4.667l-1.375-9.905h-.974V1.979h.981v-.95h-3.959v.95H11v1.033z"/><path d="M8.984 2.016H7.016V3H5.969v1.969h1.066v3.062H5.969v1.94h4.062v-1.94H8.994V4.969h1.037V3H8.984v-.984z"/></g>`),
		g.Group(children),
	)
}