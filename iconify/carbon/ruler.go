package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 10H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V11a1 1 0 0 0-1-1Zm-1 10H4v-8h4v4h2v-4h5v4h2v-4h5v4h2v-4h4Z"/>`),
		g.Group(children),
	)
}