package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsSixVerticalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108 60a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm56 16a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm-72 36a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm72 0a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm-72 68a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm72 0a16 16 0 1 0 16 16a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}