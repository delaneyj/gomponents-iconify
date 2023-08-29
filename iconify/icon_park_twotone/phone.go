package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPhone0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M8 30h32v12a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V30Z"/><path stroke-linejoin="round" d="M40 30V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v24"/><path stroke-linecap="round" d="M22 37h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPhone0)"/>`),
		g.Group(children),
	)
}