package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUUpRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 136V96H88a48 48 0 0 0 0 96h88a8 8 0 0 1 0 16H88a64 64 0 0 1 0-128h80V40a8 8 0 0 1 13.66-5.66l48 48a8 8 0 0 1 0 11.32l-48 48A8 8 0 0 1 168 136Z"/>`),
		g.Group(children),
	)
}