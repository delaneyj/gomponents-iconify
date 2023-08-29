package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallCraftAdvisory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M9.81 24.6V7.45h1.03V24.6H9.81zm1.73-9.74V7.45l8.65 3.69l-8.65 3.72z"/>`),
		g.Group(children),
	)
}