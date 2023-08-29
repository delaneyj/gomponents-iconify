package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 28V6l-4-4l-4 4v22H2v2h28v-2Zm-6 0V6.828l2-2l2 2V28Z"/>`),
		g.Group(children),
	)
}