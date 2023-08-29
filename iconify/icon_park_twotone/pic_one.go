package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PicOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPicOne0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="24" r="20" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 37l8-9l16 13"/><circle cx="18" cy="17" r="4" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m24 33l8-10l10 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPicOne0)"/>`),
		g.Group(children),
	)
}