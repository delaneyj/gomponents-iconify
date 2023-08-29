package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M36 17L12 17"/><path d="M36 31L12 31"/><path d="M17 36L17 12"/><path d="M31 36L31 12"/></g>`),
		g.Group(children),
	)
}