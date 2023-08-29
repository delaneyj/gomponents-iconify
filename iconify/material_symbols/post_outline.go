package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm15-4H6v1.5h12V17ZM6 15.5h12V14H6v1.5ZM6 12h12V6H6v6Zm0 5v1.5V17Zm0-1.5V14v1.5ZM6 12V6v6Zm0 2v-2v2Zm0 3v-1.5V17Z"/>`),
		g.Group(children),
	)
}