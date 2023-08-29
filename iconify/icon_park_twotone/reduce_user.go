package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReduceUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTReduceUser0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="12" r="8" fill="#555"/><path d="M42 44c0-9.941-8.059-18-18-18S6 34.059 6 44m13-5h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTReduceUser0)"/>`),
		g.Group(children),
	)
}