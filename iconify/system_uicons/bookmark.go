package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 4.5h6a1 1 0 0 1 1 1v12l-4-4l-4 4v-12a1 1 0 0 1 1-1z"/>`),
		g.Group(children),
	)
}