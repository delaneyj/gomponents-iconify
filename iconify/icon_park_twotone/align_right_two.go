package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignRightTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 7h24v6H16zM8 21h32v6H8zm14 14h18v6H22z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignRightTwo0)"/>`),
		g.Group(children),
	)
}