package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUUpRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 88l-48 48V40Z" opacity=".2"/><path d="M172.94 143.39a8 8 0 0 0 8.72-1.73l48-48a8 8 0 0 0 0-11.32l-48-48A8 8 0 0 0 168 40v40H88a64 64 0 0 0 0 128h88a8 8 0 0 0 0-16H88a48 48 0 0 1 0-96h80v40a8 8 0 0 0 4.94 7.39ZM184 59.31L212.69 88L184 116.69Z"/></g>`),
		g.Group(children),
	)
}