package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M30 33C30 26.5556 30 16.8889 30 4H18C18 16.8749 18 26.531 18 32.9684"/><path stroke-linecap="round" d="M30 14H25"/><path stroke-linecap="round" d="M30 20H25"/><path fill="#2F88FF" d="M18 33C18 34.0441 18 35.6103 18 37.6985C18 41.1787 20.6863 44 24 44C27.3137 44 30 41.1787 30 37.6985C30 35.6386 30 34.0937 30 33.0638L18 33Z"/></g>`),
		g.Group(children),
	)
}