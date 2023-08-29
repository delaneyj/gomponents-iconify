package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ladder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 42C10 42 10 16 10 10C10 4 16 4 16 4"/><path d="M10 14H28"/><path d="M10 22H28"/><path d="M10 30H28"/><path d="M10 38H28"/><path d="M34 43C34 43 34 17 34 11C34 5 40 5 40 5"/></g>`),
		g.Group(children),
	)
}