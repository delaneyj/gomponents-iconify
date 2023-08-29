package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WholeSiteAccelerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWholeSiteAccelerator0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="22" cy="40" r="4" fill="#fff"/><circle cx="26" cy="8" r="4" fill="#fff"/><circle cx="36" cy="24" r="4" fill="#fff"/><circle cx="12" cy="24" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 24H16m7-13l-8 10"/><path d="m33 27l-8.001 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWholeSiteAccelerator0)"/>`),
		g.Group(children),
	)
}