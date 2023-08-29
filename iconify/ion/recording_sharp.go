package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 138a117.93 117.93 0 0 0-91.84 192h-72.32A118 118 0 1 0 128 374h256a118 118 0 0 0 0-236ZM54 256a74 74 0 1 1 74 74a74.09 74.09 0 0 1-74-74Zm330 74a74 74 0 1 1 74-74a74.09 74.09 0 0 1-74 74Z"/>`),
		g.Group(children),
	)
}