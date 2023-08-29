package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m490.51 187.313l-22.628-22.626l-64.284 64.284l-64.285-64.284l-22.627 22.626l64.285 64.285l-64.285 64.285l22.627 22.627l64.285-64.284l64.284 64.284l22.628-22.627l-64.285-64.285l64.285-64.285zM145.373 120H16v272h129.373l104 104H288V16h-38.627ZM128 360H48V152h80Zm128 97.373l-96-96V150.627l96-96Z"/>`),
		g.Group(children),
	)
}