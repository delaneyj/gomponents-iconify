package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneCall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.636 4A7.364 7.364 0 0 1 20 11.364M13 8a3 3 0 0 1 3 3M7 6l-2.485.621c-.89.223-1.534 1.029-1.352 1.928c1.06 5.213 7.075 11.228 12.288 12.287c.9.183 1.705-.46 1.928-1.35l.62-2.486l-3.5-2l-1.5 1.5c-2-1-4.5-3.5-5.5-5.5L9 9.5L7 6Z"/>`),
		g.Group(children),
	)
}