package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4h10v2H11V4ZM6 7v4H4V7H1l4-4l4 4H6Zm0 10h3l-4 4l-4-4h3v-4h2v4Zm5 1h10v2H11v-2Zm-2-7h12v2H9v-2Z"/>`),
		g.Group(children),
	)
}