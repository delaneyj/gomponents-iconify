package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m158 290l-30 30L0 192L128 64l30 30l-98 98zm111 0l98-98l-98-98l30-30l128 128l-128 128z"/>`),
		g.Group(children),
	)
}