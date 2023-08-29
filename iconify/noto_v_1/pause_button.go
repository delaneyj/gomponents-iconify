package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M36.57 96h18.29V32H36.57v64zm36.57-64v64h18.29V32H73.14z"/>`),
		g.Group(children),
	)
}