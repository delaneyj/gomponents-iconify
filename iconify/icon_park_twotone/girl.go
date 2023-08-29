package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Girl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGirl0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="10" r="6" fill="#555"/><path fill="#555" d="M27.23 16h-6.46L10 36h28L27.23 16Z"/><path d="M27 36v8m-6-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGirl0)"/>`),
		g.Group(children),
	)
}