package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextVerticalAlignment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28h14v2H16zm0-5h14v2H16zm-5.154 7H13L8.64 20H6.36L2 30h2.154l.8-2h5.092zm-5.092-4L7.5 21.635L9.246 26zM2 15h28v2H2zm14-8h14v2H16zm0-5h14v2H16zm-5.154 10H13L8.64 2H6.36L2 12h2.154l.8-2h5.092zM5.754 8L7.5 3.635L9.246 8z"/>`),
		g.Group(children),
	)
}