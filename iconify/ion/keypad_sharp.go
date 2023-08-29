package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeypadSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="96" height="96" x="80" y="16" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="208" y="16" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="336" y="16" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="80" y="144" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="208" y="144" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="336" y="144" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="80" y="272" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="208" y="272" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="208" y="400" fill="currentColor" rx="8" ry="8"/><rect width="96" height="96" x="336" y="272" fill="currentColor" rx="8" ry="8"/>`),
		g.Group(children),
	)
}