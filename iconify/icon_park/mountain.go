package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M18 10L32 38H4L18 10Z"/><path stroke="#000" stroke-linecap="round" d="M28 29L33.6471 22L44 38H32"/><path stroke="#fff" stroke-linecap="round" d="M12 22L24 22"/><path stroke="#000" stroke-linecap="round" d="M14 18L10 26"/><path stroke="#000" stroke-linecap="round" d="M22 18L26 26"/></g>`),
		g.Group(children),
	)
}