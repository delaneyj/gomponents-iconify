package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSquareSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 6c-.3 0-.5.2-.5.5v5H9v-5c0-.3-.2-.5-.5-.5s-.5.2-.5.5v11c0 .3.2.5.5.5s.5-.2.5-.5v-5h6v5c0 .3.2.5.5.5s.5-.2.5-.5v-11c0-.3-.2-.5-.5-.5zm4-4h-15C3.1 2 2 3.1 2 4.5v15C2 20.9 3.1 22 4.5 22h15c1.4 0 2.5-1.1 2.5-2.5v-15C22 3.1 20.9 2 19.5 2zM21 19.5c0 .8-.7 1.5-1.5 1.5h-15c-.8 0-1.5-.7-1.5-1.5v-15C3 3.7 3.7 3 4.5 3h15c.8 0 1.5.7 1.5 1.5v15z"/>`),
		g.Group(children),
	)
}