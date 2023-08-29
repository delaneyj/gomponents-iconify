package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 0h3v2.953H1zm13 0h3v2.953h-3zM9.016 15.917c-4.334 0-7.982-3.344-7.982-7.454V4.031H3.97v3.886c0 2.908 2.23 5.13 5.047 5.13c2.826 0 5.016-2.315 5.016-5.13V4.031h2.936v4.432c-.002 4.04-3.576 7.454-7.953 7.454z"/>`),
		g.Group(children),
	)
}