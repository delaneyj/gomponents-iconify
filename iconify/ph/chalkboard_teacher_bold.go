package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardTeacherBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56 84a12 12 0 0 1 12-12h120a12 12 0 0 1 12 12v88a12 12 0 0 1-24 0V96H68a12 12 0 0 1-12-12Zm180-28v144a20 20 0 0 1-20 20h-66.74a12 12 0 0 1-11.4-8.26a36 36 0 0 0-67.74 0A12 12 0 0 1 58.74 220H40a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h176a20 20 0 0 1 20 20ZM104 164a16 16 0 1 0-16-16a16 16 0 0 0 16 16ZM212 60H44v136h6.92a60.18 60.18 0 0 1 21.76-23.16a40 40 0 1 1 62.64 0A60.18 60.18 0 0 1 157.08 196H212Z"/>`),
		g.Group(children),
	)
}