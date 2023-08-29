package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 5h10v2H18zm-7.4 0H4v2h5.4l9 20H28v-2h-8.4z"/>`),
		g.Group(children),
	)
}