package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFlipVertically0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M42 24H6"/><path fill="#fff" d="m14 4l22 12H14V4Zm0 40V32h22L14 44Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFlipVertically0)"/>`),
		g.Group(children),
	)
}