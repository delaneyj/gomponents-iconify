package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dockleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.084 7.333v16.334h24.832V7.333H3.084zm8.583 3h13.25v10.335h-13.25V10.332z"/>`),
		g.Group(children),
	)
}