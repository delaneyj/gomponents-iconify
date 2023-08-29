package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 224a8 8 0 0 1-8 8A104.11 104.11 0 0 1 96 128V51.31L61.66 85.66a8 8 0 0 1-11.32-11.32l48-48a8 8 0 0 1 11.32 0l48 48a8 8 0 0 1-11.32 11.32L112 51.31V128a88.1 88.1 0 0 0 88 88a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}