package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandageOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 12v.01M12 14v.01m-1.487-7.523L12.5 4.5a4.95 4.95 0 0 1 7 7l-2.018 2.018M15.5 15.5l-4 4a4.95 4.95 0 0 1-7-7l4-4M3 3l18 18"/>`),
		g.Group(children),
	)
}