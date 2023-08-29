package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.3 17q2.075 0 3.538-1.463T17.3 12q0-2.075-1.463-3.538T12.3 7q-.55 0-1.075.113T10.2 7.45q1.35.625 2.138 1.85t.787 2.7q0 1.475-.788 2.7T10.2 16.55q.5.225 1.025.338T12.3 17Zm-.3 6.3L8.65 20H4v-4.65L.7 12L4 8.65V4h4.65L12 .7L15.35 4H20v4.65L23.3 12L20 15.35V20h-4.65L12 23.3Z"/>`),
		g.Group(children),
	)
}