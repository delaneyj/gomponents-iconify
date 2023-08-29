package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGhost0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m8 44l4-4l4 4l4-6l4 6l4-6l4 6l4-4l4 4V20c0-8.837-7.163-16-16-16S8 11.163 8 20v24Z"/><path stroke-linecap="round" d="M19 20h2m10 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGhost0)"/>`),
		g.Group(children),
	)
}