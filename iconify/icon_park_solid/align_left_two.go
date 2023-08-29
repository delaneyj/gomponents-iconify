package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignLeftTwo0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 7h24v6H8zm0 14h32v6H8zm0 14h18v6H8z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignLeftTwo0)"/>`),
		g.Group(children),
	)
}