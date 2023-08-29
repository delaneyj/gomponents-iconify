package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M368 128H144a128 128 0 0 0 0 256h224a128 128 0 0 0 0-256Zm0 224H144a96 96 0 0 1 0-192h224a96 96 0 0 1 0 192Z"/><path fill="currentColor" d="M368 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64Zm0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}