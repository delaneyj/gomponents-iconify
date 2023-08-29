package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwordOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.938 7.937L15 4h5v5l-3.928 3.055m-2.259 1.757L11 16l-4 4l-3-3l4-4l2.19-2.815M6.5 11.5l6 6M3 3l18 18"/>`),
		g.Group(children),
	)
}