package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAUnderlineThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M62.3 171.62a4 4 0 0 0 5.32-1.92L85.36 132h85.28l17.74 37.7a4 4 0 1 0 7.24-3.4l-64-136a4 4 0 0 0-7.24 0l-64 136a4 4 0 0 0 1.92 5.32ZM128 41.39L166.87 124H89.13ZM220 216a4 4 0 0 1-4 4H40a4 4 0 0 1 0-8h176a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}