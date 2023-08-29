package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gauze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="26" cy="24" r="17" fill="#2F88FF" stroke="#000"/><circle cx="26" cy="24" r="7" fill="#43CCF8" stroke="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M5 41L26 41"/></g>`),
		g.Group(children),
	)
}