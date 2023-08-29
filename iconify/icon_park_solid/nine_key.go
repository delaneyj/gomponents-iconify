package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNineKey0"><g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path stroke="#000" d="M24.5 14a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M19 29c.818 2.193 2.548 4 5.5 4c3.038 0 5.5-2.686 5.5-6v-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNineKey0)"/>`),
		g.Group(children),
	)
}