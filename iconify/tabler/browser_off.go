package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h11a1 1 0 0 1 1 1v11m-.288 3.702A1 1 0 0 1 19 20H5a1 1 0 0 1-1-1V5c0-.276.112-.526.293-.707M4 8h4m4 0h8M3 3l18 18"/>`),
		g.Group(children),
	)
}