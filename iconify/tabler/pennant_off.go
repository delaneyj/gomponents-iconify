package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PennantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 21h4m-2 0V10m0-4V3m0 1l9 4l-4.858 2.16m-2.764 1.227L10 12M3 3l18 18"/>`),
		g.Group(children),
	)
}