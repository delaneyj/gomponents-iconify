package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M208 80h264v32H208zM40 96a64 64 0 1 0 64-64a64.072 64.072 0 0 0-64 64Zm64-32a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32Zm104 176h264v32H208zm-104 80a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64Zm0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32Zm104 176h264v32H208zm-104 80a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64Zm0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32Z"/>`),
		g.Group(children),
	)
}