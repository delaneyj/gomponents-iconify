package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M5 8C5 6.89543 5.89543 6 7 6H19L24 12H41C42.1046 12 43 12.8954 43 14V40C43 41.1046 42.1046 42 41 42H7C5.89543 42 5 41.1046 5 40V8Z"/><circle cx="22" cy="26" r="6" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M27 30L32 34"/></g>`),
		g.Group(children),
	)
}