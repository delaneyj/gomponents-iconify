package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHdd0"><g fill="none"><rect width="30" height="40" x="9" y="4" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><circle cx="32" cy="10" r="2" fill="#fff"/><circle cx="16" cy="10" r="2" fill="#fff"/><circle cx="32" cy="38" r="2" fill="#fff"/><circle cx="16" cy="38" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 30a8 8 0 1 0-8-8m8 0l-6 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHdd0)"/>`),
		g.Group(children),
	)
}