package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WoolenHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWoolenHat0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="40" height="10" x="4" y="34" fill="#555" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 26v8m9-8v8m9-8v8"/><circle cx="24" cy="8" r="4" fill="#555"/><path d="M8 34c0-8.25 1-22 16-22s16 13.75 16 22"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWoolenHat0)"/>`),
		g.Group(children),
	)
}