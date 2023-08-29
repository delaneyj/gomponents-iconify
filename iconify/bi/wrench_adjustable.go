package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrenchAdjustable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M16 4.5a4.492 4.492 0 0 1-1.703 3.526L13 5l2.959-1.11c.027.2.041.403.041.61Z"/><path d="M11.5 9c.653 0 1.273-.139 1.833-.39L12 5.5L11 3l3.826-1.53A4.5 4.5 0 0 0 7.29 6.092l-6.116 5.096a2.583 2.583 0 1 0 3.638 3.638L9.908 8.71A4.49 4.49 0 0 0 11.5 9Zm-1.292-4.361l-.596.893l.809-.27a.25.25 0 0 1 .287.377l-.596.893l.809-.27l.158.475l-1.5.5a.25.25 0 0 1-.287-.376l.596-.893l-.809.27a.25.25 0 0 1-.287-.377l.596-.893l-.809.27l-.158-.475l1.5-.5a.25.25 0 0 1 .287.376ZM3 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/></g>`),
		g.Group(children),
	)
}