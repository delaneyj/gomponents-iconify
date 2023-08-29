package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func X(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m427.314 107.313l-22.628-22.626L256 233.373L107.314 84.687l-22.628 22.626L233.373 256L84.686 404.687l22.628 22.626L256 278.627l148.686 148.686l22.628-22.626L278.627 256l148.687-148.687z"/>`),
		g.Group(children),
	)
}