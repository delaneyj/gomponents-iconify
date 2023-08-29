package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m184 128l-72 72V56Z" opacity=".2"/><path d="M117.66 50.34A8 8 0 0 0 104 56v64H32a8 8 0 0 0 0 16h72v64a8 8 0 0 0 13.66 5.66l72-72a8 8 0 0 0 0-11.32ZM120 180.69V75.31L172.69 128ZM224 40v176a8 8 0 0 1-16 0V40a8 8 0 0 1 16 0Z"/></g>`),
		g.Group(children),
	)
}