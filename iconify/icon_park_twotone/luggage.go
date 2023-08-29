package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLuggage0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="32" height="26" x="8" y="14" fill="#555" stroke-linejoin="round" rx="2"/><path d="M20 23v8"/><path stroke-linejoin="round" d="M15 40v4m18-4v4"/><path d="M28 23v8"/><path stroke-linejoin="round" d="M19 4h10m-5 0v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLuggage0)"/>`),
		g.Group(children),
	)
}