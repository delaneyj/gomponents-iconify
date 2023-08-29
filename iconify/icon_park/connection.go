package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M39 34L44 39L39 44"/><path fill="#2F88FF" d="M8 12C10.2091 12 12 10.2091 12 8C12 5.79086 10.2091 4 8 4C5.79086 4 4 5.79086 4 8C4 10.2091 5.79086 12 8 12Z"/><path stroke-linecap="round" d="M12 8L20 8C22.2091 8 24 9.79086 24 12V35C24 37.2091 25.7909 39 28 39H44"/></g>`),
		g.Group(children),
	)
}