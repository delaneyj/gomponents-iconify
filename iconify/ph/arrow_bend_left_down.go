package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32a8 8 0 0 1-8 8a88.1 88.1 0 0 0-88 88v76.69l34.34-34.35a8 8 0 0 1 11.32 11.32l-48 48a8 8 0 0 1-11.32 0l-48-48a8 8 0 0 1 11.32-11.32L96 204.69V128A104.11 104.11 0 0 1 200 24a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}