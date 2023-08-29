package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeLightKite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#09234c" d="m16.235 13.059l-4.525 8.516L21.575 30l3.249-12.561l-8.589-4.38z"/><path fill="#09234c" d="m12.751 2l-7.81 13.792l5.991 5.113l4.996-8.824l9.158 4.362l1.973-7.629L12.751 2z"/>`),
		g.Group(children),
	)
}