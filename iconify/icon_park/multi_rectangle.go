package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M12 39H44V7H12V15"/><path d="M8 39H32V15H8V23"/><path fill="#2F88FF" d="M20 23H4V39H20V23Z"/></g>`),
		g.Group(children),
	)
}