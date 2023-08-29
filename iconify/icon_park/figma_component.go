package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 12L24 5L31 12L24 19L17 12Z"/><path d="M17 36L24 29L31 36L24 43L17 36Z"/><path d="M29 24L36 17L43 24L36 31L29 24Z"/><path d="M5 24L12 17L19 24L12 31L5 24Z"/></g>`),
		g.Group(children),
	)
}