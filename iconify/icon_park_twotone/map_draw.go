package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMapDraw0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M17 12L4 6v30l13 6l14-6l13 6V12L31 6l-14 6Z"/><path d="M31 6v30M17 12v30M10.5 9l6.5 3l14-6l6.5 3m-27 30l6.5 3l14-6l6.5 3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMapDraw0)"/>`),
		g.Group(children),
	)
}