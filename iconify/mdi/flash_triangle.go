package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2L1 21h22L12 2m-2 13v-5h4l-1.5 3.5h2l-3 5.5v-4H10Z"/>`),
		g.Group(children),
	)
}