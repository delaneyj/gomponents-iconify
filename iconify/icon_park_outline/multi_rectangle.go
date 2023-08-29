package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M12 39h32V7H12v8"/><path d="M8 39h24V15H8v8"/><path d="M20 23H4v16h16V23Z"/></g>`),
		g.Group(children),
	)
}