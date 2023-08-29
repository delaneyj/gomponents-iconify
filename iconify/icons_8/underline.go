package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 4v12c0 4.43 3.57 8 8 8s8-3.57 8-8V4h-2v12c0 3.37-2.63 6-6 6s-6-2.63-6-6V4H8zM6 26v2h20v-2H6z"/>`),
		g.Group(children),
	)
}