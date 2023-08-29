package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataServer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDataServer0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path stroke="#000" d="M4 32h40"/><path stroke="#000" stroke-linejoin="round" d="M10 38h1m15 0h12"/><path stroke="#fff" stroke-linejoin="round" d="M44 37V27M4 37V27"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDataServer0)"/>`),
		g.Group(children),
	)
}