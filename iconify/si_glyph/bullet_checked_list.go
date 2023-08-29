package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletCheckedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.016 9h-3V6h3v3zM.953 8h1V6.969h-1V8zm2.016 4.986H.016V10h2.953v2.986zM.954 12h1v-1h-1v1zM5 3h10.977v.976H5zm0 4h10.977v.96H5zm0 4h10.977v.914H5zM1.366 5.295L.021 3.939l.635-.635l.71.696l2.036-2l.623.636l-2.659 2.659z"/>`),
		g.Group(children),
	)
}