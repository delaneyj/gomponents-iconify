package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpForwardToInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4H2v16h11v-2H4V8l8 5l8-5v5h2V4zm-10 7L4 6h16l-8 5zm7 4l4 4l-4 4v-3h-4v-2h4v-3z"/>`),
		g.Group(children),
	)
}