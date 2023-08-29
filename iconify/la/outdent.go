package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 7v2h13V7zm0 4v2h20v-2zm22 0v10l5-5zM3 15v2h20v-2zm0 4v2h20v-2zm0 4v2h13v-2z"/>`),
		g.Group(children),
	)
}