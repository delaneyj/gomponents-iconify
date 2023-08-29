package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.505 1.505 0 0 1-.178.71A1.5 1.5 0 0 1 13.5 15h-12A1.5 1.5 0 0 1 0 13.5v-12Zm4.85 1.642l-.35-.35l-3.5 3.5V1.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5V10h-2.293L4.854 3.146l-.005-.004Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}