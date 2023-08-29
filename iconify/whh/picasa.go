package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M986 704H768V69q117 68 186.5 185.5T1024 512q0 99-38 192zM704 404L333 33Q421 0 512 0q99 0 192 38v366zM452 240L24 668Q0 592 0 512q0-144 74-265T271 60zM256 525v430Q118 875 50 731zm64 243h635q-68 117-185.5 186.5T512 1024q-99 0-192-38V768z"/>`),
		g.Group(children),
	)
}