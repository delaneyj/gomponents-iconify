package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFlipHorizontally0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6v36"/><path fill="#555" d="m4 34l12-22v22H4Zm40 0H32V12l12 22Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFlipHorizontally0)"/>`),
		g.Group(children),
	)
}