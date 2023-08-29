package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tabletennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTabletennis0"><g fill="none"><path fill="#555" stroke="#fff" stroke-miterlimit="2" stroke-width="4" d="M31.45 16.96a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="m36 34l-2-11l-11 5l-13-3M4 41.03h40"/><circle cx="9" cy="25" r="5" fill="#fff"/><circle cx="6" cy="17" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="m34 23l10 6l-15 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTabletennis0)"/>`),
		g.Group(children),
	)
}