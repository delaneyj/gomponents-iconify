package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pisces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24L44 24"/><path d="M10 7C10 7 16 16.8203 16 24C16 31.1797 10 41 10 41"/><path d="M38 7C38 7 32 16.9597 32 24C32 31.0403 38 41 38 41"/></g>`),
		g.Group(children),
	)
}