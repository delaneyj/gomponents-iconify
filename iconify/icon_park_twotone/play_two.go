package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPlayTwo0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linecap="round" rx="3"/><path d="M18.5 24v-7.794l6.75 3.897L32 24l-6.75 3.897l-6.75 3.897V24Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPlayTwo0)"/>`),
		g.Group(children),
	)
}