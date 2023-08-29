package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Device(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2h22v4h-2V4H3v13h9v2H1V2Zm13 6h10v14H14V8Zm2 2v10h6V10h-6Zm1.998 6.998h2.004v2.004h-2.004v-2.004ZM5 20h7v2H5v-2Z"/>`),
		g.Group(children),
	)
}