package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playstore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1023H64q-26 0-45-18.5T0 959V319q0-26 19-45t45-19h128V128L320 0v255h384V0l128 128v127h128q27 0 45.5 19t18.5 45v640q0 27-18.5 45.5T960 1023zM320 383v512l448-256zM384 0h256v128H384V0z"/>`),
		g.Group(children),
	)
}