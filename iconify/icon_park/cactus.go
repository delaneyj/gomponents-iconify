package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cactus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 43L40 43"/><path stroke-linecap="round" d="M6 16C6 21.5 6.5 30 16 30"/><path stroke-linecap="round" d="M42 14C42 19 42 34 32 34"/><path fill="#2F88FF" d="M24 5C19.5817 5 16 8.58172 16 13V43H32V13C32 8.58172 28.4183 5 24 5Z"/></g>`),
		g.Group(children),
	)
}