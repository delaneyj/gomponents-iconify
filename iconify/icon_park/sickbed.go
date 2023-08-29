package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sickbed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 23L40 35"/><circle cx="12" cy="16" r="3" fill="#2F88FF"/><path fill="#2F88FF" stroke-linejoin="round" d="M29 36L29.0001 31.5L19 28V36H29Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 27.5L23 14L41.3744 20.9998C42.8515 21.5625 43.6385 23.1747 43.1737 24.6855L40 35"/></g>`),
		g.Group(children),
	)
}