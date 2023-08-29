package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eventum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 1024V0h1024v1024H0zM960 64H64v896h896V64zM832 320H384v128h256v128H384v128h448v128H192V192h640v128z"/>`),
		g.Group(children),
	)
}