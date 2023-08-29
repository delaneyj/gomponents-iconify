package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFlipVertically0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M42 24H6"/><path fill="#555" d="m14 4l22 12H14V4Zm0 40V32h22L14 44Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFlipVertically0)"/>`),
		g.Group(children),
	)
}