package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLuggage0"><g fill="none" stroke-linecap="round" stroke-width="4"><rect width="32" height="26" x="8" y="14" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="2"/><path stroke="#000" d="M20 23v8"/><path stroke="#fff" stroke-linejoin="round" d="M15 40v4m18-4v4"/><path stroke="#000" d="M28 23v8"/><path stroke="#fff" stroke-linejoin="round" d="M19 4h10m-5 0v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLuggage0)"/>`),
		g.Group(children),
	)
}