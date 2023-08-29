package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lighttraceartist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.83 4.635l-2.77 2.84H6.29a1.79 1.79 0 0 0-1.79 1.79v23.22a1.8 1.8 0 0 0 1.79 1.8h35.42a1.8 1.8 0 0 0 1.79-1.8V9.265a1.79 1.79 0 0 0-1.79-1.79H30.94l-2.77-2.84Zm18.93 5.74a1.85 1.85 0 1 1-1.64 1.651a1.859 1.859 0 0 1 1.64-1.651ZM24 13.575a8.51 8.51 0 1 1-8.51 8.51h0a8.51 8.51 0 0 1 8.51-8.51Zm0 20.71l-12.112 9.08M24 34.285l-7.267 9.08M24 34.285l-2.422 9.08M24 34.285l2.422 9.08M24 34.285l7.267 9.08M24 34.285l12.112 9.08"/>`),
		g.Group(children),
	)
}