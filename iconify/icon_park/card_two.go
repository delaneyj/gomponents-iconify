package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M28 12V4L8 14V42L20 36"/><path fill="#2F88FF" d="M20 16L40 6V34L20 44V16Z"/></g>`),
		g.Group(children),
	)
}