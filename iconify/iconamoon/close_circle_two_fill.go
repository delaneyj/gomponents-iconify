package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseCircleTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a8 8 0 1 0 8 8a1 1 0 1 1 2 0c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2a1 1 0 1 1 0 2Zm7.707.293a1 1 0 0 1 0 1.414l-.793.793l.793.793a1 1 0 0 1-1.414 1.414l-.793-.793l-.793.793a1 1 0 1 1-1.414-1.414l.793-.793l-.793-.793a1 1 0 0 1 1.414-1.414l.793.793l.793-.793a1 1 0 0 1 1.414 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}