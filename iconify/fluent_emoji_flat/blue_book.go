package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M8 26h19V4a2 2 0 0 0-2-2H8v24Z"/><path fill="#D3D3D3" d="M6 27h21v2H6v-2Z"/><path fill="#0074BA" d="M6.5 2A1.5 1.5 0 0 0 5 3.5V28h1a1 1 0 0 1 1-1h1V2H6.5Z"/><path fill="#0074BA" d="M6.5 26A1.5 1.5 0 0 0 5 27.5v1A1.5 1.5 0 0 0 6.5 30h19a1.5 1.5 0 0 0 1.415-1H7a1 1 0 1 1 0-2h20v-1H6.5Z"/></g>`),
		g.Group(children),
	)
}