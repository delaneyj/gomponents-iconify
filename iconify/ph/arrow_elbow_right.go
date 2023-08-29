package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 80v72a8 8 0 0 1-16 0V99.31l-98.34 98.35a8 8 0 0 1-11.32 0l-96-96a8 8 0 0 1 11.32-11.32L120 180.69L212.69 88H160a8 8 0 0 1 0-16h72a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}