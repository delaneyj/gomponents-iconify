package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonograph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPhonograph0"><g fill="none"><rect width="36" height="36" x="6" y="6" stroke="#fff" stroke-width="4" rx="3"/><circle cx="24" cy="25" r="11" fill="#555" stroke="#fff" stroke-width="4"/><rect width="4" height="4" x="22" y="23" fill="#fff" rx="2"/><rect width="4" height="4" x="34" y="34" fill="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m28 20l8-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPhonograph0)"/>`),
		g.Group(children),
	)
}