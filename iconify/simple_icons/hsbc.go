package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hsbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m24 12.007l-5.996 5.997V5.996L24 12.007zm-5.996-6.01H6.01l5.996 6.01l5.997-6.01zM0 12.006l6.01 5.997V5.996L0 12.007zm6.01 5.997h11.994l-5.997-5.997l-5.996 5.997z"/>`),
		g.Group(children),
	)
}