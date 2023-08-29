package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 5v22h28V5H2zm2 2h24v13.906l-5.28-5.312l-.72-.72l-.72.72l-3.81 3.812l-5.75-5.812l-.72-.72l-.72.72L4 19.874V7zm20 2a2 2 0 1 0-.001 3.999A2 2 0 0 0 24 9zm-13 6.72L20.188 25H4v-2.28l7-7zm11 2l6 6V25h-4.97l-4.155-4.188L22 17.72z"/>`),
		g.Group(children),
	)
}