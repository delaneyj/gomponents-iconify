package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cssthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m832 896l-384 128L64 896L0 0h896zM128 128l12 128h436v128H151l12 128h413v192l-128 32l-128-32V576H169l23 256l256 64l256-64l64-704H128z"/>`),
		g.Group(children),
	)
}