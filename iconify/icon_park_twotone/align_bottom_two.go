package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignBottomTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 16h6v24H7zm14-8h6v32h-6zm14 14h6v18h-6z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignBottomTwo0)"/>`),
		g.Group(children),
	)
}