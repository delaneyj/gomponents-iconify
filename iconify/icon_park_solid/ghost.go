package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGhost0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m8 44l4-4l4 4l4-6l4 6l4-6l4 6l4-4l4 4V20c0-8.837-7.163-16-16-16S8 11.163 8 20v24Z"/><path stroke="#000" stroke-linecap="round" d="M19 20h2m10 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGhost0)"/>`),
		g.Group(children),
	)
}