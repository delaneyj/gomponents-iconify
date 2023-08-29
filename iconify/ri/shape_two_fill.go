package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h5v5H2V2Zm0 15h5v5H2v-5ZM17 2h5v5h-5V2Zm0 15h5v5h-5v-5ZM8 4h8v2H8V4ZM4 8h2v8H4V8Zm14 0h2v8h-2V8ZM8 18h8v2H8v-2Z"/>`),
		g.Group(children),
	)
}