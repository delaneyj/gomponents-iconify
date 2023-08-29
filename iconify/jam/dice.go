package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 0a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h12Zm0 2H4a2 2 0 0 0-1.995 1.85L2 4v12a2 2 0 0 0 1.85 1.995L4 18h12a2 2 0 0 0 1.995-1.85L18 16V4a2 2 0 0 0-1.85-1.995L16 2Zm-1 11a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-5-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4ZM5 3a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}