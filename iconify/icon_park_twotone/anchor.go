package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAnchor0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M14 8h20M14 8h20M14 40h20"/><rect width="8" height="8" x="36" y="4" fill="#555" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="4" y="4" fill="#555" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="36" y="36" fill="#555" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="4" y="36" fill="#555" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" d="M40 14v20M8 14v20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAnchor0)"/>`),
		g.Group(children),
	)
}