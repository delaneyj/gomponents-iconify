package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobiledataOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L10 12.8v4.35l1.55-1.55L13 17l-4 4l-4-4l1.45-1.4L8 17.15V10.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM16 13.15l-2-2V6.8l-1.6 1.6L11 7l4-4l4 4l-1.4 1.4L16 6.8v6.35Z"/>`),
		g.Group(children),
	)
}