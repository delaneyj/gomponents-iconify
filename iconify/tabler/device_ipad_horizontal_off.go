package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceIpadHorizontalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h12a2 2 0 0 1 2 2v12m-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2m5 13h6M3 3l18 18"/>`),
		g.Group(children),
	)
}