package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 3.5h1v10h3l8-10v-2h-5v2h2l-5 6v-6h1v-2h-5v2Z"/>`),
		g.Group(children),
	)
}