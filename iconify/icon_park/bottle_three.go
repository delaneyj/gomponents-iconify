package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.1875 10H26.8125L33 21.1806V44H15V21.1806L21.1875 10Z"/><rect width="8" height="6" x="20" y="4" fill="#2F88FF"/><rect width="6" height="12" x="21" y="23" fill="#2F88FF" rx="3"/></g>`),
		g.Group(children),
	)
}