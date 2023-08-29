package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTable0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M14 16h20m-20 8h20m-20 8h20M18 12v24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTable0)"/>`),
		g.Group(children),
	)
}