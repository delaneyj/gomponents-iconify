package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 3.594l-.72.687l-7 7l1.44 1.44L15 7.436V24h2V7.437l5.28 5.282l1.44-1.44l-7-7l-.72-.686zM7 26v2h18v-2H7z"/>`),
		g.Group(children),
	)
}