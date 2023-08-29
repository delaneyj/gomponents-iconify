package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PayCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPayCode0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 4H4v40h40V4Z"/><path stroke-linecap="round" d="M12 16v16m8-16v16m8-16v16m8-16v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPayCode0)"/>`),
		g.Group(children),
	)
}