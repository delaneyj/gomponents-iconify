package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalCenterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignVerticalCenterTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 7h24v6H12zM8 21h32v6H8zm7 14h18v6H15z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignVerticalCenterTwo0)"/>`),
		g.Group(children),
	)
}