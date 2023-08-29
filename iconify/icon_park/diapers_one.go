package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiapersOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 6H6V24C6 28.5 8.5 42.5 24 42.5C39.5 42.5 42 28 42 24V6Z"/><path d="M6 14H16"/><path d="M32 14H42"/><path d="M42 24C32 24 25 28.8 25 42"/><path d="M6 24C16 24 23 28.8 23 42"/></g>`),
		g.Group(children),
	)
}