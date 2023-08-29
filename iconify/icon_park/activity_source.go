package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivitySource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M42 5H6V13H42V5Z"/><path d="M42 20H6V28H42V20Z"/><path d="M42 35H6V43H42V35Z"/></g>`),
		g.Group(children),
	)
}