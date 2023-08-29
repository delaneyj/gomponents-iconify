package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeLightSolidity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="m20.477 2l-4.5 8h-9l4.5-8h9" opacity=".45"/><path d="M15.975 10h9l-4.5-8h-9Z" opacity=".6"/><path d="m11.477 18l4.5-8l-4.5-8l-4.5 8Z" opacity=".8"/><path d="m11.52 30l4.5-8h9l-4.5 8h-9" opacity=".45"/><path d="M16.022 22h-9l4.5 8h9Z" opacity=".6"/><path d="m20.52 14l-4.5 8l4.5 8l4.5-8Z" opacity=".8"/>`),
		g.Group(children),
	)
}