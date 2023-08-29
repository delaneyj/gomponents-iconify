package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 168v32a56 56 0 0 1 0 112v32a88 88 0 0 0 0-176Zm-174.627-48H16v272h129.373l104 104H288V16h-38.627ZM128 360H48V152h80Zm128 97.373l-96-96V150.627l96-96Z"/>`),
		g.Group(children),
	)
}