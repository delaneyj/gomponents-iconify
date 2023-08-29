package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMergeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 224a63.66 63.66 0 0 0-37.95 12.5L160 153.36v-2a64 64 0 1 0-64 0v209.25a64 64 0 1 0 64 0V223.46l160.41 71.69A64 64 0 1 0 384 224ZM128 64a32 32 0 1 1-32 32a32 32 0 0 1 32-32Zm0 384a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm256-128a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}