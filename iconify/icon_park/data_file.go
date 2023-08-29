package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 44V4H31L40 14.5V44H8Z"/><rect width="6" height="7" x="15" y="28" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" d="M14 35H34"/><rect width="6" height="12" x="21" y="23" fill="#43CCF8" stroke="#fff"/><rect width="6" height="17" x="27" y="18" fill="#43CCF8" stroke="#fff"/></g>`),
		g.Group(children),
	)
}