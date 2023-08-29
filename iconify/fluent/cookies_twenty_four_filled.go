package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookiesTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2c.714 0 1.419.075 2.106.222a.75.75 0 0 1 .374 1.263a2.501 2.501 0 0 0 1.206 4.201a.75.75 0 0 1 .577.811a2.5 2.5 0 0 0 4.36 1.908a.75.75 0 0 1 1.307.409c.047.39.07.787.07 1.186c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm3 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-7-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm4-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM7 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}