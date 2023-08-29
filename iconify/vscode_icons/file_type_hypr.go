package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeHypr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#656868" d="M2 29.99h27.91L8.7 8.52l-4.48 8.53L2 29.99z"/><path fill="#2b2b2b" d="m2 30l17.34-10.72l-3.3-3.34L2 30z"/><path fill="#ffce01" d="M2 13.18L13.19 2H30L2 29.99V13.18z"/>`),
		g.Group(children),
	)
}