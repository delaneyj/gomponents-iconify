package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M23.9999 9L40.4544 37.5H7.54541L23.9999 9Z"/><path fill="#2F88FF" d="M23.9999 9L40.4544 37.5H7.54541L23.9999 9Z"/><path d="M4 22L19.5 17"/><path d="M28 16L44 13"/><path d="M30 19.5L44 21"/><path d="M32.7002 24L44.0002 29"/></g>`),
		g.Group(children),
	)
}