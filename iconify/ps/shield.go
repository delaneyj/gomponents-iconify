package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 0H64Q37 0 18.5 18.5T0 64v151q0 65 26.5 124T102 439l90 79l90-79q102-91 102-224V64q0-27-18.5-45.5T320 0zm21 215q0 119-87 192l-62 56V43h128q21 0 21 21v151z"/>`),
		g.Group(children),
	)
}