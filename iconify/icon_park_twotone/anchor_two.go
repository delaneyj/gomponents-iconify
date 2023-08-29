package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAnchorTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 24H4c0 11.046 8.954 20 20 20s20-8.954 20-20h-6M24 44V14"/><path fill="#555" fill-rule="evenodd" d="M24 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAnchorTwo0)"/>`),
		g.Group(children),
	)
}