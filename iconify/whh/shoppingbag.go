package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shoppingbag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1023H64q-27 0-45.5-18.5T0 959V319q0-26 18.5-45T64 255h128V127L320 0v255h384V0l128 127v128h128q26 0 45 19t19 45v640q0 27-18.5 45.5T960 1023zM384 0h256v127H384V0z"/>`),
		g.Group(children),
	)
}