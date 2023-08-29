package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 6C8 4.89543 8.89543 4 10 4H38C39.1046 4 40 4.89543 40 6V42C40 43.1046 39.1046 44 38 44H10C8.89543 44 8 43.1046 8 42V6Z"/><path stroke="#fff" stroke-linecap="round" d="M16 4V44"/><path stroke="#fff" stroke-linecap="round" d="M24 12H32"/><path stroke="#fff" stroke-linecap="round" d="M24 20H32"/><path stroke="#000" stroke-linecap="round" d="M10 4H22"/><path stroke="#000" stroke-linecap="round" d="M10 44H22"/></g>`),
		g.Group(children),
	)
}