package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M177.08 114.46A48 48 0 0 0 140 36H72a12 12 0 0 0-12 12v152a12 12 0 0 0 12 12h80a52 52 0 0 0 25.08-97.54ZM84 60h56a24 24 0 0 1 0 48H84Zm68 128H84v-56h68a28 28 0 0 1 0 56Z"/>`),
		g.Group(children),
	)
}