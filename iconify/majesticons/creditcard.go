package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creditcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 7a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1H2V7zm0 3v7a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7H2zm5 2a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}