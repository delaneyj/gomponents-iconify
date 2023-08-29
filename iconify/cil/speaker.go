package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 16H128a32.036 32.036 0 0 0-32 32v416a32.036 32.036 0 0 0 32 32h256a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32Zm0 448H128V48h256l.021 416Z"/><path fill="currentColor" d="M256 240a96 96 0 1 0 96 96a96.108 96.108 0 0 0-96-96Zm0 160a64 64 0 1 1 64-64a64.072 64.072 0 0 1-64 64Zm0-200a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64Zm0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32Z"/><path fill="currentColor" d="M240 320h32v32h-32z"/>`),
		g.Group(children),
	)
}