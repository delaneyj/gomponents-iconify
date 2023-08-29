package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChangeDateSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 5V30.0036H42V5"/><path fill="#2F88FF" d="M15.0001 23H19.0016L32.8514 8.85714L28.9943 5L15 19.143L15.0001 23Z"/><path stroke-linecap="round" d="M30 37L24 43L18 37"/><path stroke-linecap="round" d="M24 30V43"/></g>`),
		g.Group(children),
	)
}