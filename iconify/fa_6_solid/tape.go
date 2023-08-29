package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M380.8 416c41.5-40.7 67.2-97.3 67.2-160c0-123.7-100.3-224-224-224S0 132.3 0 256s100.3 224 224 224h320c17.7 0 32-14.3 32-32s-14.3-32-32-32H380.8zM224 160a96 96 0 1 1 0 192a96 96 0 1 1 0-192zm64 96a64 64 0 1 0-128 0a64 64 0 1 0 128 0z"/>`),
		g.Group(children),
	)
}