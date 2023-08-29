package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneForwardToInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18H4V8l8 5l8-5v5h-2c-2.76 0-5 2.24-5 5zm-1-7L4 6h16l-8 5z" opacity=".3"/><path fill="currentColor" d="M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h9v-2H4V8l8 5l8-5v5h2V6c0-1.1-.9-2-2-2zm-8 7L4 6h16l-8 5zm7 4l4 4l-4 4v-3h-4v-2h4v-3z"/>`),
		g.Group(children),
	)
}