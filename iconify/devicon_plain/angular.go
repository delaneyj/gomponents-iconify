package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M55.297 69.324h17.406L64 48.383ZM64 15.36l-47.668 17l7.27 63.028L64 117.762l40.398-22.375l7.27-63.028Zm29.746 78.133h-11.11l-5.99-14.953h-25.29l-5.992 14.953h-11.11L64 26.676Zm0 0"/>`),
		g.Group(children),
	)
}