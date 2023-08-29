package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOfflineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M4.281 4.225l14.496 14.496a.75.75 0 1 1-1.06 1.06l-3.786-3.785H6.693a3.687 3.687 0 0 1-3.692-3.682a3.687 3.687 0 0 1 3.568-3.68L3.22 5.285a.75.75 0 1 1 1.06-1.06zm7.72-.224c3.168 0 4.965 2.097 5.227 4.63h.08A3.687 3.687 0 0 1 21 12.314a3.686 3.686 0 0 1-2.91 3.6L7.931 5.752c.884-1.066 2.25-1.751 4.068-1.751z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}