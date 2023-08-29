package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.566 8l3.045-3.044c.42-.421.42-1.103 0-1.522L12.566.389a1.078 1.078 0 0 0-1.523 0L7.999 3.433L4.955.389a1.078 1.078 0 0 0-1.523 0L.388 3.434a1.074 1.074 0 0 0-.001 1.522L3.431 8L.387 11.044a1.075 1.075 0 0 0 .001 1.523l3.044 3.044c.42.421 1.102.421 1.523 0l3.044-3.044l3.044 3.044a1.076 1.076 0 0 0 1.523 0l3.045-3.044c.42-.421.42-1.103 0-1.523L12.566 8z"/>`),
		g.Group(children),
	)
}