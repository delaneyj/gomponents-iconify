package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v11h-2V3H3v13h9v2H1V1Zm17.25 14a2.75 2.75 0 0 1 1.947 4.693l-.01.008A2.75 2.75 0 1 1 18.25 15Zm3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412l-1.672-1.677ZM2.25 21H12v2H2.25v-2Z"/>`),
		g.Group(children),
	)
}