package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSubway0"><g fill="none"><rect width="32" height="26" x="8" y="6" stroke="#fff" stroke-width="4" rx="2"/><circle cx="14" cy="27" r="2" fill="#fff"/><circle cx="34" cy="27" r="2" fill="#fff"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 12h20v10H14z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m32 32l8 9m-23-9l-9 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSubway0)"/>`),
		g.Group(children),
	)
}