package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M41 4H7C5.34315 4 4 5.34315 4 7V41C4 42.6569 5.34315 44 7 44H41C42.6569 44 44 42.6569 44 41V7C44 5.34315 42.6569 4 41 4Z"/><path fill="#43CCF8" stroke="#fff" stroke-linejoin="round" d="M34 4V22H15V4H34Z"/><path stroke="#fff" stroke-linecap="round" d="M29 11V15"/><path stroke="#000" stroke-linecap="round" d="M11.9968 4H36.9984"/></g>`),
		g.Group(children),
	)
}