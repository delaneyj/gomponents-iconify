package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignHorizontally0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M7 17h34v14H7z"/><path stroke-linecap="round" d="M24 6v36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignHorizontally0)"/>`),
		g.Group(children),
	)
}