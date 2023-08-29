package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13 1.5H6.5a4 4 0 1 0 0 8h1m3 4.5V1.5M7.5 14V1.5"/>`),
		g.Group(children),
	)
}