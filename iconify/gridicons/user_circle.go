package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 18.5a8.5 8.5 0 1 1 .001-17.001A8.5 8.5 0 0 1 12 20.5zm0-8c-3.038 0-5.5 1.728-5.5 3.5s2.462 3.5 5.5 3.5s5.5-1.728 5.5-3.5s-2.462-3.5-5.5-3.5zm0-.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/>`),
		g.Group(children),
	)
}