package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsSixVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M102 60a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm62 10a10 10 0 1 0-10-10a10 10 0 0 0 10 10Zm-72 48a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm72 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm-72 68a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm72 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Z"/>`),
		g.Group(children),
	)
}