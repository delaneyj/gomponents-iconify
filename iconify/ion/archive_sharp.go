package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="448" height="80" x="32" y="48" fill="currentColor" rx="12" ry="12"/><path fill="currentColor" d="M64 160v280a24 24 0 0 0 24 24h336a24 24 0 0 0 24-24V160Zm192 230.63L169.32 304L192 281.32l48 48.05V208h32v121.37l48.07-48.07l22.61 22.64Z"/>`),
		g.Group(children),
	)
}