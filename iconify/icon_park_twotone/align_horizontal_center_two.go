package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignHorizontalCenterTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 12h6v24H7zm14-4h6v32h-6zm14 7h6v18h-6z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignHorizontalCenterTwo0)"/>`),
		g.Group(children),
	)
}