package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.825 5H21v6.5H9.825V5ZM21 12.5V19h-5.075v-6.5H21Zm-11.175 0h5.1V19h-5.1v-6.5ZM3 19V5h5.825v14H3Z"/>`),
		g.Group(children),
	)
}