package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowUpRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m208 80l-48 48V32Z" opacity=".2"/><path d="m213.66 74.34l-48-48A8 8 0 0 0 152 32v40H64a8 8 0 0 0-8 8v144a8 8 0 0 0 16 0V88h80v40a8 8 0 0 0 13.66 5.66l48-48a8 8 0 0 0 0-11.32ZM168 108.69V51.31L196.69 80Z"/></g>`),
		g.Group(children),
	)
}