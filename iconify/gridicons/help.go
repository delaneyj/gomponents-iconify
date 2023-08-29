package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm1 16h-2v-2h2v2zm0-4.141V15h-2v-2a1 1 0 0 1 1-1c1.103 0 2-.897 2-2s-.897-2-2-2s-2 .897-2 2H8a4 4 0 0 1 8 0a3.991 3.991 0 0 1-3 3.859z"/>`),
		g.Group(children),
	)
}