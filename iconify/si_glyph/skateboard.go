package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skateboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 5)"><path d="M15.252.268c-.548.953-1.697.857-2.978.857H3.699c-1.316 0-2.431.068-2.992-.938C.481-.219-.206.001.042.546c.649 1.42 2.046 2.402 3.663 2.402h8.573c1.573.001 2.938-.928 3.61-2.287c.284-.572-.392-.815-.636-.393z"/><ellipse cx="11.984" cy="4.947" rx=".984" ry=".947"/><ellipse cx="3.97" cy="4.959" rx=".97" ry=".959"/></g>`),
		g.Group(children),
	)
}