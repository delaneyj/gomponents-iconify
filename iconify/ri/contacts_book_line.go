package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsBookLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h16.005C20.107 2 21 2.898 21 3.99v16.02c0 1.099-.893 1.99-1.995 1.99H3V2Zm4 2H5v16h2V4Zm2 16h10V4H9v16Zm2-4a3 3 0 1 1 6 0h-6Zm3-4a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm8-6h2v4h-2V6Zm0 6h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}