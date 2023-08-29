package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 248v128q0 27-19 45.5T960 440H64q-26 0-45-18.5T0 376V248q0-26 18.5-45T64 184h896q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}