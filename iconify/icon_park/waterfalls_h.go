package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M20 6H6V17H20V6Z"/><path d="M42 31H28V42H42V31Z"/><path d="M42 6H28V23H42V6Z"/><path d="M20 25H6V42H20V25Z"/></g>`),
		g.Group(children),
	)
}