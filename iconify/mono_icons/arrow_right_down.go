package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.293 6.293a1 1 0 0 1 1.414 0L16 14.586V9a1 1 0 1 1 2 0v8a1 1 0 0 1-1 1H9a1 1 0 1 1 0-2h5.586L6.293 7.707a1 1 0 0 1 0-1.414z"/>`),
		g.Group(children),
	)
}