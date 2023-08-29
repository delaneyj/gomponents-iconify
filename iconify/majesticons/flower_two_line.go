package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13c-4 0-5-3.333-5-5V4l2.5 2L12 3l2.5 3L17 4v4c0 1.667-1 5-5 5zm0 0v8m1 0c5.6 0 7-4.667 7-7c-5.6 0-7 4.667-7 7zm0 0h-1m-1 0c-5.6 0-7-4.667-7-7c5.6 0 7 4.667 7 7zm0 0h1"/>`),
		g.Group(children),
	)
}