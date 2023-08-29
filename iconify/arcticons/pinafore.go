package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinafore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.443 42.483a17.318 17.318 0 0 1-7.571-10.834l34.256.016a17.594 17.594 0 0 1-7.563 10.814m-19.091.021h19.054M22.291 31.649V5.5M6.977 28.868H19.59V9.833ZM25.166 7.283s15.047 4.945 15.048 21.584h-15.04Z"/>`),
		g.Group(children),
	)
}