package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeEnsime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="m22.116 11.554l-6.14 3.218l-6.14-3.218l6.14-3.168Z"/><path fill="#232f3f" d="m15.975 2l12.281 6.361v6.411L15.975 8.361l-12.28 6.411V8.361L9.86 5.168L15.975 2Z"/><path fill="#f1626b" d="M15.975 14.772L3.695 8.411v6.361l12.281 6.411l12.281-6.411V8.411Z"/><path fill="#ef4551" d="M22.116 20.421L9.811 26.782L16 30l12.305-6.386Z"/><path fill="#f1626b" d="m3.695 17.228l12.281 6.361V30L3.695 23.589Z"/>`),
		g.Group(children),
	)
}