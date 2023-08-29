package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v2h16v-2H4zM18 3v8a6 6 0 1 1-12 0V3h3v8c0 1.654 1.346 3 3 3s3-1.346 3-3V3h3z"/>`),
		g.Group(children),
	)
}