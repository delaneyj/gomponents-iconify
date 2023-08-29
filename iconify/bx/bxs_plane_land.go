package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsPlaneLand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M18.842 15.296a1.61 1.61 0 0 0 1.892-1.189v-.001a1.609 1.609 0 0 0-1.177-1.949l-4.576-1.133L9.825 4.21l-2.224-.225l2.931 6.589l-4.449-.449l-2.312-3.829l-1.38.31l1.24 5.52l15.211 3.17zM3 18h18v2H3z" fill="currentColor"/>`),
		g.Group(children),
	)
}