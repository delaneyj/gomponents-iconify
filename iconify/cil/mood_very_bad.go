package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodVeryBad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16Zm147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125Z"/><path fill="currentColor" d="M256 280a104 104 0 0 0-104 104h208a104 104 0 0 0-104-104Zm-138.63-92.04l21.261-23.917l72 64l-21.26 23.918zm178.672 45.411l64-72l23.918 21.26l-64 72z"/>`),
		g.Group(children),
	)
}