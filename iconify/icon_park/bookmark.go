package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M10 44C8.89543 44 8 43.1046 8 42V6C8 4.89543 8.89543 4 10 4H38C39.1046 4 40 4.89543 40 6V42C40 43.1046 39.1046 44 38 44H10Z"/><path fill="#43CCF8" fill-rule="evenodd" stroke="#fff" stroke-linecap="round" d="M21 22V4H33V22L27 15.7273L21 22Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M10 4H38"/></g>`),
		g.Group(children),
	)
}