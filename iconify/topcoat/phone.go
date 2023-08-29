package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M31.5 6.65c0-2.589-.561-3.15-3-3.15h-17c-2.47 0-3 .529-3 3.15V36.5c0 2.439.56 3 3 3h17c2.5 0 3-.561 3-3V6.65zm-18 1.85h13v24h-13v-24zm8 28h-3v-2h3v2z"/>`),
		g.Group(children),
	)
}