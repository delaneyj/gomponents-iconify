package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinkingProblem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M38 21L43 30L38 31V37H35L29 36L28 43H13L11 32.619C7.92077 29.7028 6 25.5757 6 21C6 12.1634 13.1634 5 22 5C30.8366 5 38 12.1634 38 21Z"/><path stroke="#fff" d="M17 19C17 16.2386 19.2386 14 22 14C24.7614 14 27 16.2386 27 19C27 21.7614 24.7614 24 22 24V27"/><path stroke="#fff" d="M22 33V34"/></g>`),
		g.Group(children),
	)
}