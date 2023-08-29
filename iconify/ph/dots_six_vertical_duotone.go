package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsSixVerticalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 32v192a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16V32a16 16 0 0 1 16-16h128a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M104 60a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm60 12a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm-72 44a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm72 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm-72 68a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm72 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/></g>`),
		g.Group(children),
	)
}