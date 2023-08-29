package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19h4L17.5 8.5a2.828 2.828 0 1 0-4-4L3 15v4m9.5-13.5l4 4m-12 4l4 4M21 15v4h-8l4-4z"/>`),
		g.Group(children),
	)
}