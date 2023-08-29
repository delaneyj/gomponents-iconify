package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCircleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3a7 7 0 0 0-5.288 11.587l4.236-4.168a1.5 1.5 0 0 1 2.104 0l4.236 4.168A7 7 0 0 0 10 3Zm0 14a6.973 6.973 0 0 0 4.58-1.706l-4.23-4.163a.5.5 0 0 0-.7 0l-4.23 4.163A6.973 6.973 0 0 0 10 17Zm-8-7a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm11-2.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm1 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}