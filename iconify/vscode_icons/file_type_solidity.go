package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSolidity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#c1c1c1" d="m20.477 2l-4.5 8h-9l4.5-8h9M11.52 30l4.5-8h9l-4.5 8h-9" opacity=".45"/><path fill="#c1c1c1" d="M15.975 10h9l-4.5-8h-9Zm.047 12h-9l4.5 8h9Z" opacity=".6"/><path fill="#c1c1c1" d="m11.477 18l4.5-8l-4.5-8l-4.5 8Zm9.043-4l-4.5 8l4.5 8l4.5-8Z" opacity=".8"/>`),
		g.Group(children),
	)
}