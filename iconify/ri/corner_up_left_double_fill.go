package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeftDoubleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 19v-9h-5.586V4.586L8 11l6.414 6.414V12H18v7h2ZM10.164 6.05L8.75 4.636L2.386 11l6.364 6.364l1.414-1.414L5.214 11l4.95-4.95Z"/>`),
		g.Group(children),
	)
}