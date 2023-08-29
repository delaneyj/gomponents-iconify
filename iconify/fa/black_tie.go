package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackTie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M0 0h1536v1536H0V0zm1085 1115L864 484l221-297H451l221 297l-221 631l317 304z"/>`),
		g.Group(children),
	)
}