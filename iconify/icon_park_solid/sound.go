package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSound0"><g fill="none"><rect width="32" height="40" x="8" y="4" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><circle cx="24" cy="15" r="5" fill="#fff" stroke="#fff" stroke-width="4"/><circle cx="14" cy="10" r="2" fill="#fff"/><circle cx="14" cy="38" r="2" fill="#fff"/><circle cx="34" cy="10" r="2" fill="#fff"/><circle cx="34" cy="38" r="2" fill="#fff"/><circle cx="24" cy="32" r="6" fill="#fff" stroke="#fff" stroke-width="4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSound0)"/>`),
		g.Group(children),
	)
}