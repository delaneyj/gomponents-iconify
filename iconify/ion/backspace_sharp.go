package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 96L32 256l112 160h304V96Zm215.3 226.34L336.67 345l-65-65l-65 65L184 322.34l65-65l-65-65l22.63-22.63l65 65l65-65l22.63 22.63l-65 65Z"/>`),
		g.Group(children),
	)
}