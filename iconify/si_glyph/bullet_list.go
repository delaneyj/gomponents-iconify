package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.016 9h-3V6h3v3zM.953 8h1V6.969h-1V8zm2.016 4.986H.016V10h2.953v2.986zM.954 12h1v-1h-1v1zm2.062-7h-3V2h3v3zM.953 4h1V2.969h-1V4zM5 7h10.977v.96H5zm0-4h10.977v.96H5zm0 8h10.977v.914H5z"/>`),
		g.Group(children),
	)
}