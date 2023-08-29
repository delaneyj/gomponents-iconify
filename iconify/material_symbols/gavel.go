package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2h12v2H4Zm5.65-4.85L4 10.5l2.1-2.15L11.8 14l-2.15 2.15ZM16 9.8l-5.65-5.7L12.5 2l5.65 5.65L16 9.8ZM20.6 20L7.55 6.95l1.4-1.4L22 18.6L20.6 20Z"/>`),
		g.Group(children),
	)
}