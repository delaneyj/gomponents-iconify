package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16.042c-132.548 0-240 107.451-240 240s107.452 240 240 240s240-107.452 240-240s-107.452-240-240-240ZM403.078 403.12A207.253 207.253 0 1 1 447.667 337a207.364 207.364 0 0 1-44.589 66.12Z"/><path fill="currentColor" d="m272.614 164.987l-22.628-22.627l-113.681 113.681l113.681 113.681l22.628-22.627l-75.054-75.054H385v-32H197.56l75.054-75.054z"/>`),
		g.Group(children),
	)
}