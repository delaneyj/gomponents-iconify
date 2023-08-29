package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.364 15.535l-1.414-1.414l1.414-1.414a5 5 0 0 0-7.07-7.071L9.878 7.05L8.465 5.636l1.414-1.414a7 7 0 0 1 9.9 9.9l-1.415 1.413Zm-2.828 2.829l-1.414 1.414a7 7 0 1 1-9.9-9.9l1.414-1.414l1.415 1.414l-1.415 1.415a5 5 0 0 0 7.071 7.07l1.415-1.413l1.414 1.414Zm-.707-10.607l1.414 1.414l-7.071 7.072l-1.414-1.415l7.07-7.07Z"/>`),
		g.Group(children),
	)
}