package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16Zm147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125Z"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm-64 88a88.1 88.1 0 0 0-88 88h32a56 56 0 0 1 112 0h32a88.1 88.1 0 0 0-88-88Z"/>`),
		g.Group(children),
	)
}