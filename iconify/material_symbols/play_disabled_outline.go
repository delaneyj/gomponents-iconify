package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayDisabledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.45 13.6l-1.4-1.45l-.95-.9L8 5.15V5l11 7l-2.55 1.6Zm3.3 9L13 15.8L8 19v-8.2L1.4 4.2l1.4-1.4l18.4 18.4l-1.45 1.4ZM10 12.8Zm0 2.55l1.55-1L10 12.8v2.55Zm4.1-4.1Z"/>`),
		g.Group(children),
	)
}