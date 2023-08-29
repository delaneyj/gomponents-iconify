package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delicious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FFF" d="M0 0h128v128H0z"/><path fill="#333" d="M0 128h128v128H0z"/><path fill="#EEE" d="M128 128h128v128H128z"/><path fill="#39F" d="M128 0h128v128H128z"/>`),
		g.Group(children),
	)
}