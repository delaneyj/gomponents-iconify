package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inspection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15 17.18l-2.59-2.59L11 16l4 4l7-7l-1.41-1.41L15 17.18z"/><path fill="currentColor" d="M30 15h-2.05A12.007 12.007 0 0 0 17 4.05V2h-2v2.05A12.007 12.007 0 0 0 4.05 15H2v2h2.05A12.007 12.007 0 0 0 15 27.95V30h2v-2.05A12.007 12.007 0 0 0 27.95 17H30ZM17 25.95V23h-2v2.95A10.017 10.017 0 0 1 6.05 17H9v-2H6.05A10.017 10.017 0 0 1 15 6.05V9h2V6.05A10.017 10.017 0 0 1 25.95 15H23v2h2.949A10.017 10.017 0 0 1 17 25.95Z"/>`),
		g.Group(children),
	)
}