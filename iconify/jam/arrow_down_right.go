package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.9 7.485V1.9a1 1 0 0 1 2 0v8a.997.997 0 0 1-1 1h-8a1 1 0 1 1 0-2h5.585L.707 2.121A1 1 0 0 1 2.121.707L8.9 7.485z"/>`),
		g.Group(children),
	)
}