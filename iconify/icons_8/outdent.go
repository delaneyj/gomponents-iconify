package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 7v2h13V7H3zm0 4v2h20v-2H3zm22 0v10l5-5l-5-5zM3 15v2h20v-2H3zm0 4v2h20v-2H3zm0 4v2h13v-2H3z"/>`),
		g.Group(children),
	)
}