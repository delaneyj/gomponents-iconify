package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpToTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 14L6 24l1.4 1.4l8.6-8.6l8.6 8.6L26 24zM4 8h24v2H4z"/>`),
		g.Group(children),
	)
}