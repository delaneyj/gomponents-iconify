package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTextFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 9v11.992A1 1 0 0 1 20.007 22H3.993A.993.993 0 0 1 3 21.008V2.992C3 2.455 3.447 2 3.998 2H14v6a1 1 0 0 0 1 1h6Zm0-2h-5V2.003L21 7ZM8 7v2h3V7H8Zm0 4v2h8v-2H8Zm0 4v2h8v-2H8Z"/>`),
		g.Group(children),
	)
}