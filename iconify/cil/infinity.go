package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M463.2 176.805a112 112 0 0 0-158.39 0l-48.57 48.568l-48.573-48.573a112 112 0 1 0 0 158.392l48.568-48.569l48.57 48.569A112 112 0 1 0 463.2 176.805ZM185.04 312.569a80 80 0 1 1 0-113.138l55.2 55.2v2.746Zm255.528 0a80 80 0 0 1-113.136 0l-55.2-55.2v-2.744l55.2-55.2a80 80 0 1 1 113.136 113.144Z"/>`),
		g.Group(children),
	)
}