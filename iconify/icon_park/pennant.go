package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pennant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 44H16M12 6V4V6ZM12 22V44V22ZM12 44H8H12Z"/><path d="M8 44H12H16"/><path fill="#2F88FF" d="M12 6V22L40 14L12 6Z"/></g>`),
		g.Group(children),
	)
}