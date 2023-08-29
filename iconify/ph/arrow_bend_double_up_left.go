package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDoubleUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M85.66 146.34a8 8 0 0 1-11.32 11.32l-48-48a8 8 0 0 1 0-11.32l48-48a8 8 0 0 1 11.32 11.32L43.31 104ZM128 96H99.31l34.35-34.34a8 8 0 0 0-11.32-11.32l-48 48a8 8 0 0 0 0 11.32l48 48a8 8 0 0 0 11.32-11.32L99.31 112H128a88.1 88.1 0 0 1 88 88a8 8 0 0 0 16 0A104.11 104.11 0 0 0 128 96Z"/>`),
		g.Group(children),
	)
}