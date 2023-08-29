package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataQualityDefinition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.6 24.4l2.6 2.6l-2.6 2.6L26 31l4-4l-4-4zm-2.2 0L19.8 27l2.6 2.6L21 31l-4-4l4-4z"/><circle cx="11" cy="8" r="1" fill="currentColor"/><circle cx="11" cy="16" r="1" fill="currentColor"/><circle cx="11" cy="24" r="1" fill="currentColor"/><path fill="currentColor" d="M24 3H8c-1.1 0-2 .9-2 2v22c0 1.1.9 2 2 2h7v-2H8v-6h18V5c0-1.1-.9-2-2-2zm0 16H8v-6h16v6zm0-8H8V5h16v6z"/>`),
		g.Group(children),
	)
}