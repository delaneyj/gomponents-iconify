package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 48v416h480V48Zm448 384H48v-96h416Zm0-128H48v-96h416ZM48 176V80h416v96Z"/><path fill="currentColor" d="M80 112h32v32H80zm0 128h32v32H80zm0 128h32v32H80z"/>`),
		g.Group(children),
	)
}