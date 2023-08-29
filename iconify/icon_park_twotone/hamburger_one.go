package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HamburgerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHamburgerOne0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 22v-1c0-8.837-7.163-16-16-16S8 12.163 8 21v1m0 12s-2 5 2 7s24 2 28 0s2-7 2-7"/><rect width="38" height="14" x="5" y="21" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="7"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m11 29l1.592-.796A5.996 5.996 0 0 1 17.5 28l.424.17A4.979 4.979 0 0 0 22 28v0a4.979 4.979 0 0 1 4.076-.17l.96.384a4.71 4.71 0 0 0 3.857-.16v0a4.712 4.712 0 0 1 4.214 0L37 29"/><circle cx="17" cy="15.05" r="2.5" fill="#fff"/><circle cx="23.75" cy="12.3" r="2.5" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHamburgerOne0)"/>`),
		g.Group(children),
	)
}