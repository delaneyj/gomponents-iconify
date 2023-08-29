package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMobileOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.159 3.185C7.415 3.066 7.699 3 8 3h8a2 2 0 0 1 2 2v9m0 4v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V6m5-2h2M3 3l18 18m-9-4v.01"/>`),
		g.Group(children),
	)
}