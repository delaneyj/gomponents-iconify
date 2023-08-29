package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRecord0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="38" height="24" x="5" y="18" fill="#fff" stroke="#fff" stroke-linecap="round" rx="2"/><path stroke="#fff" stroke-linecap="round" d="M8 12h32M15 6h18"/><path stroke="#000" stroke-linecap="round" d="M26 24v6"/><path stroke="#000" d="M18 32.75c0-1.52 1.29-2.75 2.88-2.75H26v3.25c0 1.52-1.29 2.75-2.88 2.75h-2.24C19.29 36 18 34.77 18 33.25v-.5Z"/><path stroke="#000" stroke-linecap="round" d="m31 25l-5-1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRecord0)"/>`),
		g.Group(children),
	)
}