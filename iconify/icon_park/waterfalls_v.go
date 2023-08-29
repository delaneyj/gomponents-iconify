package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M23 20L23 6L6 6L6 20L23 20Z"/><path d="M42 42V28L25 28L25 42H42Z"/><path d="M31 6V20H42V6H31Z"/><path d="M6 28L6 42H17V28H6Z"/></g>`),
		g.Group(children),
	)
}