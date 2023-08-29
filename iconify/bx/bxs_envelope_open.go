package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsEnvelopeOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M21.555 8.168l-9-6a1 1 0 0 0-1.109 0l-9 6A.995.995 0 0 0 2.004 9H2v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9h-.004a.997.997 0 0 0-.441-.832zM20 12.7L12 17l-8-4.3v-2.403l8 4.299l8-4.299V12.7z" fill="currentColor"/>`),
		g.Group(children),
	)
}