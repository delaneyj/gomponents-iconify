package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trident(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 6l2-2v3a7 7 0 0 0 14 0V4l2 2"/><path d="M12 21V3l-2 2m4 0l-2-2"/></g>`),
		g.Group(children),
	)
}