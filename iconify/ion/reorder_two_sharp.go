package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReorderTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="44" d="M118 304h276m-276-96h276"/>`),
		g.Group(children),
	)
}