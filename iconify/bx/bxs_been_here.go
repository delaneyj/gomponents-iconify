package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsBeenHere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8zm-1 13.414l-3.707-3.707l1.414-1.414L11 12.586l5.293-5.293l1.414 1.414L11 15.414z" fill="currentColor"/>`),
		g.Group(children),
	)
}