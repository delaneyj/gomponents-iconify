package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonFillTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 2.5a5.25 5.25 0 0 0-2.519 9.857a9.005 9.005 0 0 0-6.477 8.37a.75.75 0 0 0 .727.773H20.27a.75.75 0 0 0 .727-.772a9.005 9.005 0 0 0-6.477-8.37A5.25 5.25 0 0 0 12 2.5Z"/>`),
		g.Group(children),
	)
}