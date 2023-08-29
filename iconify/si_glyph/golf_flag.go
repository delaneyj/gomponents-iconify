package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="5.433" cy="14.98" rx="5.433" ry=".98"/><path d="M5 0h.956v12.959H5zm2.031.103s1.644-.42 4.45 1.461c2.806 1.882 2.249 3.512 4.103 3.646c.998-.066-.42 1.168-1.854.604c-2.781-1.093-3.6-2.637-6.699-1.21V.103z"/><circle cx="13.945" cy="13.945" r="1.945"/></g>`),
		g.Group(children),
	)
}