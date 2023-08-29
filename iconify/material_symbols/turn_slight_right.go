package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSlightRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20v-7.575q0-.4.15-.763t.425-.637L14.6 6h-2.25V4H18v5.65h-2V7.4l-5 5V20H9Z"/>`),
		g.Group(children),
	)
}