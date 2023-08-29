package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 1024V576h448v448H576zM576 0h448v448H576V0zM0 576h448v448H0V576zM0 0h448v448H0V0z"/>`),
		g.Group(children),
	)
}