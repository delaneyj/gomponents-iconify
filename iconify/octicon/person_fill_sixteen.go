package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonFillSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.243 4.757a3.757 3.757 0 1 1 5.851 3.119a6.006 6.006 0 0 1 3.9 5.339a.75.75 0 0 1-.715.784H2.721a.75.75 0 0 1-.714-.784a6.006 6.006 0 0 1 3.9-5.34a3.753 3.753 0 0 1-1.664-3.118Z"/>`),
		g.Group(children),
	)
}