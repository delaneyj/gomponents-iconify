package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchAccessOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17v-2h2v2H3Zm0-8V7h2v2H3Zm4 12v-2h2v2H7ZM7 5V3h2v2H7Zm8 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 12v-2h2v2h-2Zm0-8V7h2v2h-2ZM7 17V7h10v10H7Zm2-2h6V9H9v6Zm3-3Z"/>`),
		g.Group(children),
	)
}